/*
只有满足这些标准的方法才会被远程调用视为可见；其他的方法都会被忽略：

- 方法是外部可见的。
- 方法有两个参数，参数的类型都是外部可见的。
- 方法的第二个参数是一个指针。
- 方法有返回类型错误*/
package main

import (
    "net/rpc"
    "net"
    "fmt"
    "net/http"
)

const (
	PROTO = "tcp"   
	PORT = ":9999"
)


type Person struct {
	Name string
	Age int
	Score float32
}

// SelfIntroduce is func to introduce self.
func (p *Person) SelfIntroduce(person Person, res *string) (error) {
	msg := "My name is %s, i am %d years old."
    *res = fmt.Sprintf(msg, person.Name, person.Age)
    fmt.Println(*res)
	return nil
}

func main() {
    person := new(Person)
    rpc.Register(person)
    rpc.HandleHTTP()
    l, e := net.Listen(PROTO,PORT)
    if e != nil {
        fmt.Println("listen error:", e)
    }
    http.Serve(l, nil)

}