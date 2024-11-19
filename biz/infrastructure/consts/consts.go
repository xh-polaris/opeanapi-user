package consts

import (
	"errors"
	"time"
)

const DefaultPageSize int64 = 10

var ErrNotAuthentication = errors.New("not authentication")
var ErrForbidden = errors.New("forbidden")
var PageSize int64 = 10

// 数据库相关
const (
	ID           = "_id"
	UserID       = "user_id"
	Status       = "status"
	DeleteStatus = 3
	EffectStatus = 0
	Content      = "content"
	TxId         = "tx_id"
)

// 默认值
const (
	DefaultExpire         = time.Hour * 24 * 30
	DefaultKeyName        = "default-"
	DefaultDeveloperName  = "开发者用户"
	DefaultEnterpriseName = "企业用户"
)

// 提示词
const (
	RemainIncrease     = "余额 "
	RemainInsufficient = "余额不足"
)
