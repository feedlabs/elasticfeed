package main

import (
	"fmt"
	"net/rpc"
	"log"
)

type Args struct {
	X, Y int
}

func main(){

	client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Synchronous call
	args := &Args{1,2}
	var reply int
	err = client.Call("bbb.Add", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Result: %d+%d=%d", args.X, args.Y, reply)
}
