package key

import (
	"errors"
	"github.com/xh-polaris/opeanapi-user/biz/infrastructure/config"
	"github.com/xh-polaris/opeanapi-user/biz/infrastructure/consts"
	util "github.com/xh-polaris/opeanapi-user/biz/infrastructure/util/page"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/basic"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"time"
)

const (
	prefixKeyCacheKey = "cache:key"
	CollectionName    = "key"
)

type IMongoMapper interface {
	Insert(ctx context.Context, k *Key) error
	Update(ctx context.Context, k *Key) error
	UpdateWithTime(ctx context.Context, k *Key, updateTime time.Time) error
	FindOne(ctx context.Context, id string) (*Key, error)
	Delete(ctx context.Context, id string) error
	FindAndCount(ctx context.Context, userId string, p *basic.PaginationOptions) ([]*Key, int64, error)
	Count(ctx context.Context, userId string) (int64, error)
}

type MongoMapper struct {
	conn *monc.Model
}

func NewMongoMapper(config *config.Config) *MongoMapper {
	conn := monc.MustNewModel(config.Mongo.URL, config.Mongo.DB, CollectionName, config.Cache)
	return &MongoMapper{
		conn: conn,
	}
}

func (m *MongoMapper) Insert(ctx context.Context, k *Key) error {
	if k.ID.IsZero() {
		k.ID = primitive.NewObjectID()
		k.CreateTime = time.Now()
		k.UpdateTime = time.Now()
	}
	key := prefixKeyCacheKey + k.ID.Hex()
	_, err := m.conn.InsertOne(ctx, key, k)
	return err
}

func (m *MongoMapper) Update(ctx context.Context, k *Key) error {
	k.UpdateTime = time.Now()
	key := prefixKeyCacheKey + k.ID.Hex()
	_, err := m.conn.UpdateByID(ctx, key, k.ID, bson.M{"$set": k})
	return err
}

func (m *MongoMapper) UpdateWithTime(ctx context.Context, k *Key, updateTime time.Time) error {
	k.UpdateTime = updateTime
	key := prefixKeyCacheKey + k.ID.Hex()
	_, err := m.conn.UpdateByID(ctx, key, k.ID, bson.M{"$set": k})
	return err
}

func (m *MongoMapper) FindOne(ctx context.Context, id string) (*Key, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, consts.ErrInvalidObjectId
	}
	var k Key
	key := prefixKeyCacheKey + oid.Hex()
	err = m.conn.FindOne(ctx, key, &k,
		bson.M{
			consts.ID:     oid,
			consts.Status: bson.M{"$ne": consts.DeleteStatus},
		})

	switch {
	case err == nil:
		return &k, nil
	case errors.Is(err, monc.ErrNotFound):
		return nil, consts.ErrNotFound
	default:
		return nil, err
	}
}

func (m *MongoMapper) FindAndCount(ctx context.Context, userId string, p *basic.PaginationOptions) (keys []*Key, total int64, err error) {
	skip, limit := util.ParsePageOpt(p)
	keys = make([]*Key, 0, limit)
	err = m.conn.Find(ctx, &keys,
		bson.M{ // 根据userid查找未删除的key
			consts.UserID: userId,
			consts.Status: bson.M{"$ne": consts.DeleteStatus}},
		&options.FindOptions{
			Skip:  &skip,
			Limit: &limit,
			Sort:  bson.M{consts.ID: 1},
		})
	if err != nil {
		return nil, 0, err
	}

	total, err = m.conn.CountDocuments(ctx, bson.M{
		consts.UserID: userId,
		consts.Status: bson.M{"$ne": consts.DeleteStatus},
	})
	if err != nil {
		return nil, 0, err
	}
	return keys, total, nil
}

func (m *MongoMapper) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return consts.ErrInvalidObjectId
	}
	var k Key
	key := prefixKeyCacheKey + oid.Hex()
	err = m.conn.FindOne(ctx, key, &k, bson.M{consts.ID: oid})

	if err != nil {
		return err
	}

	k.DeleteTime = time.Now()
	k.UpdateTime = time.Now()
	k.Status = consts.DeleteStatus
	_, err = m.conn.UpdateByID(ctx, key, k.ID, bson.M{"$set": k})
	return err
}

func (m *MongoMapper) Count(ctx context.Context, userId string) (int64, error) {
	total, err := m.conn.CountDocuments(ctx, bson.M{consts.UserID: userId})
	if err != nil {
		return 0, err
	}
	return total, nil
}
