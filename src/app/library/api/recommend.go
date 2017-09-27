package api

import (
	_ "encoding/json"
	_ "errors"
	"fmt"
)

type Recommend struct {
}

func NewRecommend() *Recommend {
	return &Recommend{}
}

func (this *Recommend) InitUserFollow(uid int) {
	url := ""
	iclient := New()
	response, err := iclient.Get(url)
	fmt.Println(response, err)
}
