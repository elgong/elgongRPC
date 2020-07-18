package server

import (
	"reflect"
	"sync"
)

type service struct {
	name    string
	typ     reflect.Type
	rcvr    reflect.Value
	methods sync.Map
}
