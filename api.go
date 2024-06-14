package main

import "net/http"

type Foo struct {
}

func (f *Foo) Sum(args interface{}, r *http.Request) any {
	x := make(map[string]int32)
	x["hello"] = 1
	return x
}

func (f *Foo) Multiply(args struct {
	x int32
	y int32
}, r *http.Request) any {

	return struct {
		result int32
	}{
		args.x * args.y,
	}
}
