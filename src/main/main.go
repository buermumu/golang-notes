package main

import (
	_ "app"
	_ "fmt"
	"github.com/buermumu/mcq"
	//_ "github.com/go-sql-driver/mysql"
	"net"
)

/***
unique request id usage
*/

func main() {
	//var c chan int
	//c = make(chan int, 10)

	dns := "127.0.0.1:11212"
	m, err := mcq.New()
	addr, err := net.ResolveTCPAddr("tcp", dns)
	if err != nil {
		panic(err)
	}
	m.Get(addr, "user_recommend_articles")
}
