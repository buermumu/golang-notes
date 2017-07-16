package base

import (
	"fmt"
	"net/http"
	"strings"
)

// 保存请求上下文, 对请求进行格式化
type Conext struct {
	Res http.ResponseWriter
	Req *http.Request
}

func (t *Conext) Run(res http.ResponseWriter, req *http.Request) {
	t.Res = res
	t.Req = req
	t.Parse()
}

func (t *Conext) Parse() {
}

// 完成controller的公共方法
type Request struct {
	cxt    Conext
	Data   map[string]string
	Querys map[string]string
}

func (t *Request) Init(w http.ResponseWriter, r *http.Request) {
	t.cxt.Run(w, r)
	t.Data = make(map[string]string)
}

func (t *Request) Gw() http.ResponseWriter {
	return t.cxt.Res
}

func (t *Request) Gr() *http.Request {
	return t.cxt.Req
}

func (t *Request) Gq(k string) string {
	qes := t.cxt.Req.URL.Query()
	str := strings.Join(qes[k], "")
	return str
}

func (t *Request) Q(k string) string {
	fmt.Println(t.Data)
	return "bbb"
}
