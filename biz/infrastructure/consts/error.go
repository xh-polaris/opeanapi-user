package consts

import "errors"

var (
	ErrInvalidObjectId = errors.New("invalid objectId")
	ErrNotFound        = errors.New("not found")
)
