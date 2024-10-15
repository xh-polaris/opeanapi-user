package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username   string             `bson:"username" json:"username"`
	Role       int                `bson:"role" json:"role"`
	Auth       bool               `bson:"auth" json:"auth"`
	AuthId     string             `bson:"auth_id" json:"authId"`
	Remain     int64              `bson:"remain" json:"remain"`
	Status     int                `bson:"status" json:"status"`
	CreateTime time.Time          `bson:"create_time,omitempty" json:"createTime"`
	UpdateTime time.Time          `bson:"update_time,omitempty" json:"updateTime"`
	DeleteTime time.Time          `bson:"delete_time,omitempty" json:"deleteTime"`
}
