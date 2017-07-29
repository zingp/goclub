package main

import (
	"math/rand"
	"fmt"
	"time"
)

/*随机数、抽奖、验证码、负载均衡【随机或者轮询】*/
func main()  {
	//种子，否则每次生成的随机数会是一样的，可以放在init函数中。
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		// 生成int伪随机数
		a := rand.Int()
		fmt.Println(a)
	}

	for i := 0; i < 10; i++ {
		// 循环10次，每次产生[0,100)的随机整数
		a := rand.Intn(100)
		fmt.Println(a)
	}

	for i := 0; i < 10; i++ {
		// (0,1)之间的浮点数
		a := rand.Float32()
		fmt.Println(a)
	}
}
