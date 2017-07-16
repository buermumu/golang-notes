package app

import (
	"flag"
	"log"
	"net/http"
)

import (
	"app/routing"
)

func init() {

}

func initCofig() {

}

func initDebug() {

}

func ListenHttp() {
	var addr = flag.String("addr", ":1799", "127.0.0.1")
	err := http.ListenAndServe(*addr, routing.Default{})
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
