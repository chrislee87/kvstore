package kvstore

import (
	"errors"
)

var (
	ErrOutOfCapacity = errors.New("Out of capacity, can't store more data!")
	ErrDataNotFound  = errors.New("No matching date!")
)
