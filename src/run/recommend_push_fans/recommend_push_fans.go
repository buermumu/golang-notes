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
	var task_worker = 20
	task_list := make(chan map[string]string, task_worker)
	waiting := make(chan int)

	mclient, err := mcq.New()
	if err != nil {
		panic(err)
	}
	// read task data write to task_list
	go func(task_list chan<- map[string]string, mclient *mcq.Client) {
		for {
			item, err := read(mclient)
			if err != nil {
				error_log(err)
				panic(err)
			}
			if item != nil {
				task_list <- item
			}
			time.Sleep(500 * time.Millisecond)
		}
	}(task_list, mclient)

	// Bug : 这里起来N多协程， 系统被拖死， 改为只起10个， 然后每个协程中for{}获取，得到就处理
	for i := 0; i < task_worker; i++ {
		// process task
		go func(task_list <-chan map[string]string) {
			for {
				select {
				case value := <-task_list:
					handler(value)
				}
			}
		}(task_list)
	}

	// 阻塞
	<-waiting
}

func handler(item map[string]string) {
	uid := item["uid"]
	rid := item["rid"]
	fans_list, err := getFans(uid)
	if err != nil {
		error_log(err)
		panic(err)
	}
	for _, fuid := range fans_list {
		last_id := insert(fuid, rid)
		debug_log(fmt.Sprintf("last_id:%d uid:%s rid:%s", last_id, fuid, rid))
	}
	msg := fmt.Sprintf("%s, %s done", item["uid"], item["rid"])
	fmt.Println(msg)
}

func read(clent *mcq.Client) (map[string]string, error) {
	dns := "127.0.0.1:11212"
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
	res, err := stmt.Exec(rid, uid, time.Now().Unix())
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
