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
	uri := fmt.Sprintf("%s/follower/fans?uid=%s&count=%d", domain, uid, count)
	result, err := this.getApiValue(uri)
	if err != nil {
		return nil, err
	}
	if result["err_code"].(float64) <= 0 {
		return nil, errors.New(fmt.Sprintf("%s", result["err_msg"]))
	}
	var uids []string
	for _, uid := range result["data"].([]interface{}) {
		uids = append(uids, uid.(string))
	}
	return uids, nil
}

func (this *Follower) GetFollows(uid string, count int) ([]string, error) {
	uri := fmt.Sprintf("%s/follower/friends?uid=%s&count=%d", domain, uid, count)
	result, err := this.getApiValue(uri)
	if err != nil {
		return nil, err
	}
	if result["err_code"].(float64) <= 0 {
		return nil, errors.New(fmt.Sprintf("%s", result["err_msg"]))
	}
	var uids []string
	for _, uid := range result["data"].([]interface{}) {
		uids = append(uids, uid.(string))
	}
	return uids, nil
}

func (this *Follower) getApiValue(url string) (map[string]interface{}, error) {
	iclient := New()
	response, err := iclient.Get(url)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	json.Unmarshal(response, &result)
	return result, nil
}
