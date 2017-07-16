package mcq

import (
	"bytes"
	"fmt"
	"math/rand"
	"net"
	"sync"
	"time"
)

var (
	DefaultMaxIdleConns int = 5
)

type resource struct {
	conn net.Conn
}

type Client struct {
	mu           sync.Mutex
	MaxIdleConns int // 最大空闲链接
	freeConn     []*resource
}

func New() (*Client, error) {
	return &Client{}, nil
}

func (c *Client) dial(addr net.Addr) (*resource, error) {
	conn, err := net.Dial(addr.Network(), addr.String())
	if err != nil {
		return nil, err
	}
	r := &resource{
		conn: conn,
	}
	return r, err
}

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

func (c *Client) getMaxIdleConns() int {
	if c.MaxIdleConns == 0 {
		c.MaxIdleConns = DefaultMaxIdleConns
	}
	return c.MaxIdleConns
}

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
	return r, err
}

func (c *Client) GetCount() int {
	return len(c.freeConn)
}

func (c *Client) Set(addr net.Addr, k string, value string) {
	r, err := c.getConn(addr)
	if err != nil {
		panic(err)
	}
	var cmd bytes.Buffer
	cmd.WriteString(fmt.Sprintf("set %s %s\r\n", k, value))
	n, err := r.conn.Write(cmd.Bytes())
	fmt.Println(n, err)
}

func (c *Client) GetItem(addr net.Addr, i int) {
	r, err := c.getConn(addr)
	if err != nil {
		panic(err)
	}
	num := rand.Intn(10)
	time.Sleep(time.Duration(num) * time.Second)
	c.releaseFreeConn(r)
}
