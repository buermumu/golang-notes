package user

/**
1.init 完成map[string]controlerStruct 注册
2.分析url进行modules, contrller, action的注册
*/

import (
	"app"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
)

type Profile struct {
	response http.ResponseWriter // 该类中操作响应体
	request  *http.Request       // 该类中获取请求
	test     int
}

func (t *Profile) Handle(response http.ResponseWriter, request *http.Request) {
	t.response = response
	t.request = request
	t.test = 1

	// 处理业务逻辑
	//t.getData()

	md5ctx := md5.New()
	md5ctx.Write([]byte("frank"))
	val := md5ctx.Sum(nil)

	// 输出数据给客户端
	io.WriteString(t.response, "this is user.profile response data")
	fmt.Println(hex.EncodeToString(val))
}

func (t *Profile) Result(response http.ResponseWriter, request *http.Request) {
	t.response = response
	t.request = request
	io.WriteString(t.response, fmt.Sprintf("%s%s", t.test, "abc"))
}

func (t *Profile) getData() {
	fmt.Println("this is gettest")
}

func init() {
	fmt.Println("user.profile init")
	app.Register("user.profile", &Profile{})
}
