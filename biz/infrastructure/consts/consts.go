package consts

import (
	"errors"
)

const DefaultPageSize int64 = 10

var ErrNotAuthentication = errors.New("not authentication")
var ErrForbidden = errors.New("forbidden")
var PageSize int64 = 10
