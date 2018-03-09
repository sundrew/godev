package redis

import (
	"sync"
	"time"
)

var debug = true

const (
	DefaultMaxConnNumber    = 100
	DefaultMaxIdleNumber    = 25
	DefaultMaxIdleSeconds   = 28
	DefaultMaxConnWaitTimes = 50
)

// connection pool of only one redis
type Pool struct {
	Address  string
	Password string

	// 统计信息
	IdleNum         int
	ActiveNum       int
	MaxConnNum      int
	MaxIdleNum      int
	CreateNum       int
	CreateFailedNum int
	WaitTimeoutNum  int
	PingErrNum      int
	CallNetErrNum   int
	MaxIdleSeconds  int64

	ClientPool chan *Conn
	mu         sync.RWMutex

	CallNum int64
	callMu  sync.RWMutex

	ScriptMap   map[string]string
	CallConsume map[string]int // 命令消耗时长
}

// include multi redis server's connection pool
type MultiPool struct {
	pools map[string]*Pool
}

func Now() string {
	return time.Now().Format("2006-01-02 15:04:05 ")
}

func Debug(info, address string) {
	if debug {
		println(Now() + info + "|addr=" + address)
	}
}
