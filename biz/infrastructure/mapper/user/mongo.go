package user

import (
	"context"
	"errors"
	"github.com/xh-polaris/openapi-user/biz/infrastructure/config"
	"github.com/xh-polaris/openapi-user/biz/infrastructure/consts"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	prefixUserCacheKey = "cache:user"
	CollectionName     = "user"
)

type IMongoMapper interface {
	Insert(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	FindOne(ctx context.Context, id string) (*User, error)
	Delete(ctx context.Context, id string) error
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

func (m *MongoMapper) Insert(ctx context.Context, user *User) error {
	if user.ID.IsZero() {
		user.ID = primitive.NewObjectID()
		user.CreateTime = time.Now()
		user.UpdateTime = user.CreateTime
	}
	key := prefixUserCacheKey + user.ID.Hex()
	_, err := m.conn.InsertOne(ctx, key, user)
	return err
}

func (m *MongoMapper) Update(ctx context.Context, user *User) error {
	user.UpdateTime = time.Now()
	key := prefixUserCacheKey + user.ID.Hex()
	_, err := m.conn.UpdateByID(ctx, key, user.ID, bson.M{"$set": user})
	return err
}

func (m *MongoMapper) FindOne(ctx context.Context, id string) (*User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, consts.ErrInvalidObjectId
	}
	var user User
	key := prefixUserCacheKey + oid.Hex()
	err = m.conn.FindOne(ctx, key, &user, bson.M{
		consts.ID:     oid,
		consts.Status: bson.M{"$ne": consts.DeleteStatus},
	})

	switch {
	case err == nil:
		return &user, nil
	case errors.Is(err, monc.ErrNotFound):
		return nil, consts.ErrNotFound
	default:
		return nil, err
	}
}

func (m *MongoMapper) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return consts.ErrInvalidObjectId
	}
	var user User
	key := prefixUserCacheKey + oid.Hex()
	err = m.conn.FindOne(ctx, key, &user, bson.M{consts.ID: oid})

	if err != nil {
		return err
	}

	user.DeleteTime = time.Now()
	user.UpdateTime = time.Now()
	user.Status = consts.DeleteStatus
	_, err = m.conn.UpdateByID(ctx, key, user.ID, bson.M{"$set": user})
	return err
}
