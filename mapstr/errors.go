package mapstr

import "errors"

var (
	ErrKeyNotFound       = errors.New("key not found in path")
	ErrUnableToCastToMap = errors.New("unable to cast to Mapstr")
)
