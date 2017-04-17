package kvstore

import (
	"errors"
)

var (
	ErrOutOfCapacity = errors.New("Out of capacity, can't store more data!")
	ErrDateLesser    = errors.New("Appending with a lesser date!")
	ErrDateNotFound  = errors.New("No matching date!")
)
