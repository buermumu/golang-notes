package main

import (
	_ "app"
	_ "fmt"
	"github.com/buermumu/mcq"
	_ "github.com/go-sql-driver/mysql"
	"net"
)

/***
unique request id usage
*/

func main() {
	//var c chan int
	//c = make(chan int, 10)

	dns := "127.0.0.1:6379"
	m, err := mcq.New()
	addr, err := net.ResolveTCPAddr("tcp", dns)
	if err != nil {
		panic(err)
	}
	m.Set(addr, "fa", "this is value")
	/*
		if err != nil {
			panic(err)
		}
		for i := 0; i < 10; i++ {
			go func(dns string, c chan<- int, i int, m *mcq.Client) {
				m.GetItem(addr, i)
				c <- i
			}(dns, c, i, m)
		}

		for i := 0; i < 10; i++ {
			select {
			case <-c:
			}
		}
		fmt.Println("last", m, m.GetCount())
	*/
	//app.ListenHttp()

}
