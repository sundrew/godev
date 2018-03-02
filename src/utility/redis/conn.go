package redis

import "time"

const (
	ConnectionTimeout = 10e9
	ReadTimeout       = 60e9
	WriteTimeout      = 60e9
	DefaultBufferSize = 64

	RetryWaitSeconds = time.Second
	RetryTimes       = 2

	TypeError        = '-'
	TypeSimpleString = '+'
	TypeBulkString   = '$'
	TypeIntegers     = ':'
	TypeArrays       = '*'
)

