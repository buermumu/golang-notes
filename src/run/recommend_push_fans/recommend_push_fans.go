package main

import (
	"app/library/api"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/buermumu/mcq"
	_ "github.com/go-sql-driver/mysql"
	"net"
	"os"
	"time"
)

func main() {
	task_ch := make(chan map[string]string, 20)
	go func(task_ch chan<- map[string]string) {
		for {
			item, err := read()
			if err != nil {
				error_log(err)
				panic(err)
			}
			if item == nil {
				return
			}
			task_ch <- item
		}
	}(task_ch)

	for {
		select {
		case v := <-task_ch:
			fmt.Println(v)
		}
	}

}

func process() {
	handler(item["uid"], item["rid"])
}

func handler(uid, rid string) {
	fans_list, err := getFans(uid)
	if err != nil {
		error_log(err)
		panic(err)
	}
	for _, fuid := range fans_list {
		last_id := insert(fuid, rid)
		debug_log(fmt.Sprintf("last_id:%d uid:%s rid:%s", last_id, fuid, rid))
	}
}

func read() (map[string]string, error) {
	dns := "127.0.0.1:11212"
	mcq, err := mcq.New()
	addr, err := net.ResolveTCPAddr("tcp", dns)
	if err != nil {
		error_log(err)
		panic(err)
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
		error_log(err)
		panic(err)
	}
	res, err := stmt.Exec(uid, rid, time.Now().Unix())
	if err != nil {
		error_log(err)
		panic(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		error_log(err)
		panic(err)
	}
	return id
}

func error_log(err error) {
	message := fmt.Sprintf("%s", err)
	filename := "error.error_log"
	error_log_file := fmt.Sprintf("%s/%s", "./", filename)
	f, err := os.OpenFile(error_log_file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(message)
	f.WriteString("\n")
}

func debug_log(message string) {
	filename := "debug_log"
	error_log_file := fmt.Sprintf("%s/%s", "./", filename)
	f, err := os.OpenFile(error_log_file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(message)
	f.WriteString("\n")
}
