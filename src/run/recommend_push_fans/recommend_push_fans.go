package main

import (
	"app/library/api"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/buermumu/mcq"
	_ "github.com/go-sql-driver/mysql"
	"net"
	"time"
)

/***
unique request id usage
*/

func main() {
	process()
}

func process() {
	item, err := read()
	if err != nil {
		panic(err)
	}
	if item == nil {
		return
	}
	fans_list, err := getFans(item["uid"])
	for _, uid := range fans_list {
		last_id := insert(uid, item["rid"])
		fmt.Println("last_id:%s uid:%s rid:%s", last_id, uid, item["rid"])
	}
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

func getFans(uid string) ([]string, error) {
	f := api.NewFollower()
	list, err := f.GetFans(uid, 100)
	return list, err
}

func insert(uid string, rid string) int64 {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/geekbook?charset=utf8")
	stmt, err := db.Prepare(`INSERT gk_recommend_feed (rid, uid, create_time) values (? , ?, ?)`)
	if err != nil {
		panic(err)
	}
	res, err := stmt.Exec(uid, rid, time.Now().Unix())
	if err != nil {
		panic(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	return id
}
