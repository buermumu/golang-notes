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
	_ "time"
)

func main() {
	mclient, err := mcq.New()
	if err != nil {
		panic(err)
	}

	uid, err := read(mclient)
	fmt.Println(uid, err)

	/**
	task_num := 10
	waiting := make(chan int)
	task_list := make(chan int, task_num)
	go func(task_list <-chan int, mclient *mcq.Client) {
		task_list <- 12
		fmt.Println(<-task_list)
	}(task_list, mclient)
	**/
	//<-waiting
}

func read(client *mcq.Client) ([]byte, error) {
	dns := "127.0.0.1:11212"
	addr, err := net.ResolveTCPAddr("tcp", dns)
	if err != nil {
		panic(err)
	}
	uid, err := client.Get(addr, "new_register_user")
	return uid, err
}
