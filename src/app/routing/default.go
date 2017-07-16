package routing

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

import (
	_ "app/controller/mobile"
	"app/library/helper"
	"app/register"
)

type Default struct {
}

func (d Default) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var rk, action string
	rk = d.formatRouteKey(r)
	cnt, ok := register.GetController(rk)
	if !ok {
		return
	}

	action = d.getActionName(r)
	action = helper.Ucfirst(action)
	object := reflect.New(cnt.Type())
	parent := object.MethodByName("Init")
	if !parent.IsValid() {
		fmt.Println("parent not valid")
		return
	}
	args := []reflect.Value{reflect.ValueOf(w), reflect.ValueOf(r)}
	parent.Call(args)
	method := object.MethodByName(action)
	if !method.IsValid() {
		return
	}
	method.Call(nil)
}

func (d Default) formatRouteKey(r *http.Request) string {
	paths := strings.Split(r.URL.Path, "/")
	return strings.Join(paths[1:len(paths)-1], "_")
}

func (d Default) getActionName(r *http.Request) string {
	paths := strings.Split(r.URL.Path, "/")
	action := paths[len(paths)-1:]
	return action[0]
}
