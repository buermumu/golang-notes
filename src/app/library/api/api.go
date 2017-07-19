package api

import (
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

var (
	defaultTimeOut time.Duration = 30 * time.Second
	domain         string        = "http://u.api.geekbook.cc"
)

type Iclient struct {
	c         *http.Client
	timeout   time.Duration
	keepalive time.Duration
}

func New() *Iclient {
	iclient := new(Iclient)
	defaultTransport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   defaultTimeOut,
			KeepAlive: defaultTimeOut,
		}).Dial,
	}
	iclient.c = &http.Client{
		Transport: defaultTransport,
	}
	return iclient
}

func (this *Iclient) Get(uri string) (buf []byte, err error) {
	resp, err := this.c.Get(uri)
	if err != nil {
		panic(err)
	}
	buf, err = ioutil.ReadAll(resp.Body)
	return buf, err
}
