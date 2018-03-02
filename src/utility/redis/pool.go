package redis

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
	WaitTimeoutNum int
	PingErrNum int
	CallNetErrNum int
	MaxIdleSeconds int64

	ClientPool chan *Conn
}

// include multi redis server's connection pool
type MultiPool struct {
	pools map[string]*Pool
}
