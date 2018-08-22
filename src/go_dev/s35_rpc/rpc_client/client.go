package main

import (
    "net/rpc"
    "fmt"
)

const (
	PROTO = "tcp"
    PORT = "127.0.0.1:9999"
)

type Person struct {
	Name string
	Age int
	Score float32
}

func main() {
    client, err := rpc.DialHTTP(PROTO, PORT)
    if err != nil {
        fmt.Println("dialing:", err)
    }

    person := &Person{
		Name: "LiuYouyuan",
		Age: 26,
		Score: 100.0,
	}
	var ret *string
    err = client.Call("Person.SelfIntroduce", person, &ret)
    if err != nil {
         fmt.Println("arith error:", err)
	}
	
	fmt.Println(*ret)
	
	person2 := &Person{
		Name: "Jeo Chen",
		Age: 35,
		Score: 99.0,
	}
    
    var res *string
    resCall  := client.Go("Person.SelfIntroduce", person2,&res,nil) //异步调用
    replyCall := <-resCall.Done
    if replyCall.Error != nil {
        fmt.Println("arith error:", replyCall.Error)
    }
    fmt.Println(*res)
}