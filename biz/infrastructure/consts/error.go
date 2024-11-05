package consts

import "errors"

var (
	ErrInvalidObjectId    = errors.New("invalid objectId")
	ErrNotFound           = errors.New("not found")
	ErrParse              = errors.New("key parse error")
	ErrKeyUnavailable     = errors.New("key is unavailable")
	ErrHostUnavailable    = errors.New("host is unavailable")
	ErrUpdate             = errors.New("update failed")
	ErrInSufficientRemain = errors.New("insufficient remain")
)
