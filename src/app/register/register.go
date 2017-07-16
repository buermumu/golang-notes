package register

import (
	_ "fmt"
	"reflect"
	"sync"
)

var (
	rsMu sync.Mutex
)

var map_controllers = make(map[string]reflect.Value)

func AddController(k string, v reflect.Value) {
	rsMu.Lock()
	defer rsMu.Unlock()
	if k == "" {
		panic("Register controllers container args[key] is nil")
	}
	if _, ok := map_controllers[k]; ok {
		panic("Register controllers container is exists")
	}

	map_controllers[k] = v
}

func GetController(k string) (reflect.Value, bool) {
	v, ok := map_controllers[k]
	if !ok {
		var _v reflect.Value
		return _v, ok
	}
	return v, ok
}
