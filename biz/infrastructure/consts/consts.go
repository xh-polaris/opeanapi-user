package consts

import (
	"errors"
	"time"
)

const DefaultPageSize int64 = 10

var ErrNotAuthentication = errors.New("not authentication")
var ErrForbidden = errors.New("forbidden")
var PageSize int64 = 10

const (
	ID             = "_id"
	UserID         = "user_id"
	Status         = "status"
	DeleteStatus   = 3
	DefaultExpire  = time.Hour * 24 * 30
	DefaultKeyName = "default-"
)
