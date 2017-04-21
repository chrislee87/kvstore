package kvstore

import (
	"errors"
)

var (
	ErrOutOfCapacity    = errors.New("Out of capacity, can't store more data!")
	ErrDataNotFound     = errors.New("No matching date!")
	ErrCompressCodec    = errors.New("Error compress codec!")
	ErrWrongConfig      = errors.New("Wrong config!")
	ErrFailedCreateFile = errors.New("Failed to create file!")
)
