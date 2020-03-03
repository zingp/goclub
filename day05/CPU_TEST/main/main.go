package main

import (
	"math/rand"
	"fmt"
	"time"
	"regexp"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyz"

//生成指定长度的随机字符串
func RandStringBytes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func geneArray(n int) []string {
	var array []string

	for i:=0;i<n;i++ {
		st := RandStringBytes(6)
		array = append(array, st)
	}

	return array
}


//打印匹配到的字符串
func Print() {

	array := geneArray(100000)

	for i:=0; i< len(array); i++ {
		res := regexp.MustCompile(`xnf`)
		ret_array := res.FindAllString(array[i], -1)
		if len(ret_array) != 0{
			fmt.Println(array[i])
		}
	}

    fmt.Println("==========")

    for j:=0; j<len(array); j++ {
		res2 := regexp.MustCompile(`xnf`)
		ret_string := res2.FindString(array[j])
		if len(ret_string) != 0 {
			fmt.Println(array[j])
		}
	}
}

func main()  {
	for {
		Print()
		time.Sleep(time.Second * 5)
	}
}
