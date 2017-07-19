package api

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Follower struct {
}

func NewFollower() *Follower {
	f := new(Follower)
	return f
}

func (this *Follower) GetFans(uid string, count int) ([]interface{}, error) {
	iclient := New()
	uri := fmt.Sprintf("%s/follower/fans?uid=%s&count=%d", domain, uid, count)
	response, err := iclient.Get(uri)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	json.Unmarshal(response, &result)
	if result["err_code"].(float64) <= 0 {
		return nil, errors.New(fmt.Sprintf("%s", result["err_msg"]))
	}
	return result["data"].([]interface{}), nil
}
