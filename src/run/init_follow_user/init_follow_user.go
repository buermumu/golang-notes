package main

import (
	"app/library/api"
	"bytes"
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
				select {
				case uid := <-task_list:
					handler(uid)
				}
				time.Sleep(500 * time.Millisecond)
			}
		}(task_list)
	}

	fmt.Println("end")

	<-waiting
}

func read(client *mcq.Client) ([]byte, error) {
	dns := "127.0.0.1:11212"
	addr, err := net.ResolveTCPAddr("tcp", dns)
	if err != nil {
		panic(err)
	}
	uid, err := client.Get(addr, "new_register_user")
	last_uid := bytes.TrimRight(uid, "\r\n")
	return last_uid, err
}

func handler(uid string) {
	api_recommend := api.NewRecommend()
	api_recommend.InitUserFollow(uid)

}
