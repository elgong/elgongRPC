package main

import (
	"log"
	"net/rpc"
)

func main(){

	client, err := rpc.DialHTTP("tcp", httpAddr)
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call
	// 入参
	args := &Args{7, 8}
	// 返回参数
	var reply int

	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
}