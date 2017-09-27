package api

import (
	_ "encoding/json"
	_ "errors"
	"fmt"
	_ "reflect"
)

type Recommend struct {
}

func NewRecommend() *Recommend {
	return &Recommend{}
}

func (this *Recommend) InitUserFollow(uid string) {
	url := fmt.Sprintf("%s/recommend/userinit?uid=%s", domain, uid)
	iclient := New()
	response, err := iclient.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(url, "ok")
}
