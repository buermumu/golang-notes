package main

import (
	_ "app/library/api"
	_ "database/sql"
	_ "encoding/json"
	"fmt"
	"github.com/buermumu/mcq"
	_ "github.com/go-sql-driver/mysql"
	"net"
	_ "os"
	"time"
)

func main() {
	mclient, err := mcq.New()
	if err != nil {
		panic(err)
	}
	task_work := 10
	waiting := make(chan int)
	task_list := make(chan string, task_work)

	// read
	go func(task_list chan<- string, mclient *mcq.Client) {
		for {
			uid, err := read(mclient)
			if err != nil {
				panic(err)
			}
			task_list <- string(uid)
			time.Sleep(500 * time.Millisecond)
		}
	}(task_list, mclient)

	// handler
	for i := 0; i < task_work; i++ {
		go func(task_list <-chan string) {
			for {
				uid := <-task_list
				if len(uid) > 0 {
					api_recommend := api.NewRecommend()
					api_recommend.InitUserFollow(uid)
				}
			}
		}(task_list)
	}

	<-waiting
}

func read(client *mcq.Client) ([]byte, error) {
	dns := "127.0.0.1:11212"
	addr, err := net.ResolveTCPAddr("tcp", dns)
	if err != nil {
		panic(err)
	}
	uid, err := client.Get(addr, "new_register_user")
	uid = []byte{49, 48, 48, 51, 55, 54, 54, 51, 56, 49, 50, 56, 56, 48, 55, 48, 51, 13, 10}
	return uid, err
}
