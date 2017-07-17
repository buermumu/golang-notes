package mcq

import (
	"bufio"
	"fmt"
	_ "math/rand"
	"net"
	"sync"
	_ "time"
)

var (
	DefaultMaxIdleConns int = 5
)

const (
	RESPONSE_ERROR []byte = []byte("ERROR\r\n")
	RESPONSE_END   []byte = []byte("END\r\n")
	RESPONSE_VALUE []byte = []byte("VALUE")
)

type resource struct {
	conn net.Conn
	rw   *bufio.ReadWriter
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
	r.rw = bufio.NewReadWriter(bufio.NewReader(r.conn), bufio.NewWriter(r.conn))
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

func (c *Client) Get(addr net.Addr, k string) {
	r, err := c.getConn(addr)
	if err != nil {
		panic(err)
	}
	_, err = fmt.Fprintf(r.rw, "get user_recommend_articles\r\n")
	r.rw.Flush()
	if err != nil {
		panic(err)
	}
	c.parseResponse(r)
	c.releaseFreeConn(r)
}

/**
VALUE geekbook_post_article_test 0 194
{"uid":"1006299791130764","aid":"20029776248047601","url":"http:\/\/colobu.com\/2017\/06\/27\/Lint-your-golang-code-like-a-mad-man\/?hmsr=toutiao.io&utm_medium=toutiao.io&utm_source=toutiao.io"}

*/
func (c *Client) parseResponse(r *resource) {
	s := bufio.NewScanner(r.rw)
	fmt.Println(s.Text())
	/*
		for {
			result, err := r.rw.ReadBytes('\n')
			if err != nil {
				panic(err)
			}
			if result == RESPONSE_ERROR {
				fmt.Println("parse error.")
				break
			}
			if result == RESPONSE_END {
				fmt.Println("parse end.")
				break
			}
			break
		}
	*/
}
