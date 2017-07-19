package main

import (
	"app/library/api"
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
	fmt.Println(item, err)
}

func read() ([]byte, error) {
	dns := "127.0.0.1:11212"
	mcq, err := mcq.New()
	addr, err := net.ResolveTCPAddr("tcp", dns)
	if err != nil {
		return nil, err
	}
	result, err := mcq.Get(addr, "user_recommend_articles")
	return result, err
}

func getFans(uid string) ([]interface{}, error) {
	f := api.NewFollower()
	list, err := f.GetFans(uid, 100)
	return list, err
}

func insert() {

}
