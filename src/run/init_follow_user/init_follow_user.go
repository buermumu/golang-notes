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
	debug_log("begin")

	// read
	go func(task_list chan<- string, mclient *mcq.Client) {
		for {
			debug_log("process read a")
			uid, err := read(mclient)
			if err != nil {
				debug_log("process read error")
				panic(err)
			}
			task_list <- string(uid)
			time.Sleep(500 * time.Millisecond)
			debug_log("process read b")
		}
	}(task_list, mclient)

	// handler
	for i := 0; i < task_work; i++ {
		go func(task_list <-chan string) {
			for {
				debug_log("process handler a")
				select {
				case uid := <-task_list:
					handler(uid)
				}
				time.Sleep(500 * time.Millisecond)
				debug_log("process handler b")
			}
		}(task_list)
	}

	fmt.Println("end")

	<-waiting
}

func read(client *mcq.Client) ([]byte, error) {
	debug_log("read func a")
	dns := "127.0.0.1:11212"
	addr, err := net.ResolveTCPAddr("tcp", dns)
	if err != nil {
		panic(err)
	}
	uid, err := client.Get(addr, "new_register_user")
	last_uid := bytes.TrimRight(uid, "\r\n")
	debug_log("read func b")
	return last_uid, err
}

func handler(uid string) {
	if len(uid) > 0 {
		debug_log("handler func a")
		api_recommend := api.NewRecommend()
		api_recommend.InitUserFollow(uid)
		debug_log("handler func b")
	}
}

func debug_log(message string) {
	filename := "x_log"
	debug_log_file := fmt.Sprintf("%s/%s", "./", filename)
	f, err := os.OpenFile(debug_log_file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(message)
	f.WriteString("\n")
}
