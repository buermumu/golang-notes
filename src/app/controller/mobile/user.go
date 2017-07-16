package mobile

/**
1.init 完成map[string]controlerStruct 注册
2.分析url进行modules, contrller, action的注册
*/

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
)

import (
	"app/base"
	"app/register"
)

type Udata struct {
	Uid      int    `json:"uid"`
	Nickname string `json:"nickname"`
	Age      int    `json:"age"`
}

type Profile struct {
	base.Request
	v1 int
}

func (t *Profile) Profile() {
	t.Data["name"] = "Frank"
	fmt.Println(t.Gq("a"))
	x := Udata{111, "Frnak", 23}
	js, ok := json.Marshal(x)
	if ok != nil {
		fmt.Println("Parse json fial")
	}
	jstr := fmt.Sprintf("%s", js)

	var ud Udata
	err := json.Unmarshal(js, &ud)
	fmt.Println("err:", err)
	fmt.Println(ud.Nickname)

	io.WriteString(t.Gw(), fmt.Sprintf("%s", "profile.page"))
	io.WriteString(t.Gw(), jstr)
}

func init() {
	register.AddController("mobile_user", reflect.ValueOf(Profile{}))
}
