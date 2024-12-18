package key

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Key struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	UserId     string             `bson:"user_id" json:"userId"`
	Name       string             `bson:"name" json:"name"`
	Content    string             `bson:"content" json:"content"`
	Status     int                `bson:"status" json:"status"`
	Hosts      []string           `bson:"hosts" json:"hosts"`
	Timestamp  time.Time          `bson:"timestamp" json:"timestamp"`
	ExpireTime time.Time          `bson:"expire_time" json:"expireTime"`
	CreateTime time.Time          `bson:"create_time" json:"createTime"`
	UpdateTime time.Time          `bson:"update_time" json:"updateTime"`
	DeleteTime time.Time          `bson:"delete_time,omitempty" json:"deleteTime"`
}
