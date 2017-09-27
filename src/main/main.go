package main

import (
	_ "app"
	_ "app/library/api"
	ft "fmt"
	_ "github.com/buermumu/mcq"
	//_ "github.com/go-sql-driver/mysql"
	_ "net"
)

/***
unique request id usage
*/

func main() {
	//var c chan int
	//c = make(chan int, 10)

	ft.Println("test")
	/*
		return nil
		dns := "127.0.0.1:11212"
		m, err := mcq.New()
		addr, err := net.ResolveTCPAddr("tcp", dns)
		if err != nil {
			panic(err)
		}
		result, err := m.Get(addr, "user_recommend_articles")
		fmt.Println(result, err)
	*/
}
