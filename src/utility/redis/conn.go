package redis

import (
	"bufio"
	"errors"
	"net"
	"strings"
	"sync"
	"time"
)

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

var (
	ErrNil        = errors.New("nil data return")
	ErrBadType    = errors.New("invalid return type")
	ErrBadTcpConn = errors.New("invalid tcp conn")

	CommonErrPreFix = "CommonError:"
)

type Conn struct {
	sync.RWMutex
	Address        string
	keepAlive      bool
	isIdle         bool
	pipeCount      int
	lastActiveTime int64
	buffer         []byte
	conn           *net.TCPConn
	rb             *bufio.Reader
	wb             *bufio.Writer
	readTimeout    time.Duration
	writeTimeout   time.Duration
	connectTimeout time.Duration
	pool           *Pool
	err            error // 表示该条链接是否已经出错
	isOnce         bool  // 用于判断每次调用后，是否自动放回连接池，true自动放回无需开发者显式操作，默认为false
}

func NewConn(conn *net.TCPConn, connectTimeout, readTimeout, writeTimeout time.Duration, keepAlive bool, pool *Pool, Address string) *Conn {
	return &Conn{
		conn:           conn,
		lastActiveTime: time.Now().Unix(),
		keepAlive:      keepAlive,
		isIdle:         true,
		buffer:         make([]byte, DefaultBufferSize),
		rb:             bufio.NewReader(conn),
		wb:             bufio.NewWriter(conn),
		readTimeout:    readTimeout,
		writeTimeout:   writeTimeout,
		connectTimeout: connectTimeout,
		pool:           pool,
		Address:        Address,
		isOnce:         false,
	}
}

func Connect(addr string, connectTimeout, readTimeout, writeTimeout time.Duration) (*Conn, error) {
	addrPass := strings.Split(addr, ":")
	address := ""
	password := ""
	if len(addrPass) == 3 {
		address = addrPass[0] + ":" + addrPass[1]
		password = addrPass[2]
	} else if len(addrPass) == 2 {
		address = addr
	} else {
		return nil, errors.New("invalid address pattern")
	}
	return Dial(address, password, connectTimeout, readTimeout, writeTimeout, false, nil)
}

// connect with timeout
func Dial(address, password string, connectTimeout, readTimeout, writeTimeout time.Duration, keepAlive bool, pool *Pool) (*Conn, error) {
	c, e := net.DialTimeout("tcp", address, connectTimeout)
	if e != nil {
		return nil, e
	}
	if _, ok := c.(*net.TCPConn); !ok {
		return nil, ErrBadTcpConn
	}
	if password != "" {
		address = address + ":" + password
	}
	conn := NewConn(c.(*net.TCPConn), connectTimeout, readTimeout, writeTimeout, keepAlive, pool, address)
	if password != "" {
		if _, e := conn.AUTH(password); e != nil {
			return nil, e
		}
	}
	return conn, nil
}

// call redis command with request => response model
func (c *Conn) Call(command string, args ...interface{}) (interface{}, error) {
	// 如果连接已经被标记出错，直接返回
	// 应用场景，OnceConn没有获取到连接，会新建一个err非nil的Conn结构
	if c.err != nil{
		return nil, c.err
	}

	c.lastActiveTime = time.Now().Unix()
	if c.pool != nil{
		c.pool.callMu.Lock()
	}
}
