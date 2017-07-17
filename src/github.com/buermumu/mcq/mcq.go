package mcq

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	_ "math/rand"
	"net"
	"sync"
	"time"
)

// errors
var (
	cmd_error = errors.New("command error.")
)

// variables
var (
	DefaultMaxIdleConns int    = 5
	delim_error         []byte = []byte("ERROR\r\n")
	delim_end           []byte = []byte("END\r\n")
	delim_value         []byte = []byte("VALUE")
)

// Conn resource
type resource struct {
	id            int // unique int
	multiplex     int // multiplexing count
	conn          net.Conn
	rw            *bufio.ReadWriter
	borth_time    int // Birthday
	last_use_time int // last use time
	last_cmd      string
}

// Client
type Client struct {
	mu           sync.Mutex
	MaxIdleConns int // 最大空闲链接
	freeConn     []*resource
}

func New() (*Client, error) {
	return &Client{}, nil
}

// Dial
func (c *Client) dial(addr net.Addr) (*resource, error) {
	conn, err := net.Dial(addr.Network(), addr.String())
	if err != nil {
		return nil, err
	}
	t := time.Now()
	r := &resource{
		conn:          conn,
		rw:            bufio.NewReadWriter(bufio.NewReader(r.conn), bufio.NewWriter(r.conn)),
		borth_time:    t.UnixNano(),
		last_use_time: t.UnixNano(),
	}
	return r, err
}

// Put free conn to pool
func (c *Client) putFreeConn(r *resource) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if len(c.freeConn) == c.getMaxIdleConns() {
		r.conn.Close()
		return nil
	}
	c.freeConn = append(c.freeConn, r)
	return nil
}

func (c *Client) releaseFreeConn(r *resource) {
	c.putFreeConn(r)
}

// Get max free conn number
func (c *Client) getMaxIdleConns() int {
	if c.MaxIdleConns == 0 {
		c.MaxIdleConns = DefaultMaxIdleConns
	}
	return c.MaxIdleConns
}

// Get conn from poll
func (c *Client) getConn(addr net.Addr) (*resource, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if len(c.freeConn) > 0 {
		r := c.freeConn[len(c.freeConn)-1]
		c.freeConn = c.freeConn[:len(c.freeConn)-1]
		return r, nil
	}
	r, err := c.dial(addr)
	if err != nil {
		return nil, err
	}
	r.multiplex++
	return r, err
}

// Read a message
func (c *Client) Get(addr net.Addr, k string) ([]byte, error) {
	r, err := c.getConn(addr)
	if err != nil {
		panic(err)
	}
	r.last_cmd = "get user_recommend_articles\r\n"
	_, err = fmt.Fprintf(r.rw, r.last_cmd)
	r.rw.Flush()
	if err != nil {
		panic(err)
	}
	result, err := c.parseResponse(r)
	c.releaseFreeConn(r)
	return result, err
}

// Parse response data
func (c *Client) parseResponse(r *resource) ([]byte, error) {
	var buf []byte
	for {
		result, err := r.rw.ReadBytes('\n')
		if err != nil {
			return nil, err
		}
		if bytes.Equal(result, delim_error) {
			return nil, cmd_error
		}
		if bytes.Equal(result, delim_end) {
			break
		}
		if bytes.Equal(result[0:5], delim_value) {
			continue
		}
		buf = append(buf, result...)
	}
	return buf, nil
}
