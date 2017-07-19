package main

import (
	"app/library/api"
	"encoding/json"
	"fmt"
	"github.com/buermumu/mcq"
	//_ "github.com/go-sql-driver/mysql"
	"net"
)

/***
unique request id usage
*/

func main() {
	item, err := read()
	if err != nil {
		panic(err)
	}
	fans_list, err := getFans(item["uid"])
	fmt.Pirntln(fans_list, err)
}

func read() (map[string]string, error) {
	dns := "127.0.0.1:11212"
	mcq, err := mcq.New()
	addr, err := net.ResolveTCPAddr("tcp", dns)
	if err != nil {
		return nil, err
	}
	var data map[string]string
	result, err := mcq.Get(addr, "user_recommend_articles")
	json.Unmarshal(result, &data)
	return data, err
}

func getFans(uid string) ([]interface{}, error) {
	f := api.NewFollower()
	list, err := f.GetFans(uid, 100)
	return list, err
}

func insert(uid string, rid string) {

}
