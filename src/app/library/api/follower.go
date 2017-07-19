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

func (this *Follower) GetFans(uid string, count int) ([]string, error) {
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
	var uids []string
	for _, uid := range result["data"].([]interface{}) {
		uids = append(uids, uid.(string))
	}
	return uids, nil
}
