package transaction

import (
	"github.com/xh-polaris/openapi-user/biz/infrastructure/config"
	"github.com/xh-polaris/openapi-user/biz/infrastructure/consts"
	"github.com/xh-polaris/openapi-user/biz/infrastructure/mapper/account"
	"github.com/xh-polaris/openapi-user/biz/infrastructure/mapper/user"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"time"
)

const (
	UserCollectionName    = "user"
	AccountCollectionName = "account"
)

type IUserTransaction interface {
	UpdateRemain(ctx context.Context, id string, increment int64, txId string) error
}

type UserTransaction struct {
	userConn    *monc.Model
	accountConn *monc.Model
}

func NewUserTransaction(config *config.Config) *UserTransaction {
	userConn := monc.MustNewModel(config.Mongo.URL, config.Mongo.DB, UserCollectionName, config.Cache)
	accountConn := monc.MustNewModel(config.Mongo.URL, config.Mongo.DB, AccountCollectionName, config.Cache)
	return &UserTransaction{
		userConn:    userConn,
		accountConn: accountConn,
	}
}

func (t *UserTransaction) UpdateRemain(ctx context.Context, id string, increment int64, txId string) error {
	s, err := t.userConn.StartSession()
	if err != nil {
		return err
	}
	defer s.EndSession(ctx)

	_, err = s.WithTransaction(ctx, func(ctx mongo.SessionContext) (interface{}, error) {
		// 查找用户
		oid, err2 := primitive.ObjectIDFromHex(id)
		if err2 != nil {
			return nil, consts.ErrInvalidObjectId
		}
		var aUser user.User
		err2 = t.userConn.FindOneNoCache(ctx, &aUser, bson.M{
			consts.ID:     oid,
			consts.Status: bson.M{"$ne": consts.DeleteStatus},
		})
		if err2 != nil {
			return nil, consts.ErrNotFound
		}

		// 判断是否足够
		if (increment > 0) || (aUser.Remain+increment > 0) {
			// 余额足够
			_, err3 := t.userConn.UpdateByIDNoCache(ctx, aUser.ID, bson.M{
				"$inc": bson.M{
					"remain": increment,
				},
				"$set": bson.M{
					"update_time": time.Now(),
				},
			})
			if err3 != nil {
				return nil, consts.ErrUpdate
			}

			// TODO 新增流水
			aAccount := &account.Account{
				ID:         primitive.NewObjectID(),
				TxId:       txId,
				Increment:  increment,
				UserId:     id,
				CreateTime: time.Now(),
			}
			_, err4 := t.accountConn.InsertOneNoCache(ctx, aAccount)
			if err4 != nil {
				return nil, consts.ErrAccount
			}

			return aUser, nil
		}
		// 余额不足够
		return aUser, consts.ErrInSufficientRemain
	})

	return err
}
