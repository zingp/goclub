package main

import (
	"fmt"
	"sort"
)

func mapTest() {
	a := make(map[int]int)
	a[0] = 10
	a[1] = 11
	a[2] = 12
	a[3] = 13
	//for k,v := range a {
	//	fmt.Println(k, "==", v)
	//}     // map遍历是无序的

	var s []int
	for k,_ := range a {
		s = append(s, k)
	}

	sort.Ints(s)    //排序

	for _, key := range s {
		fmt.Println(a[key])
	}
}

func main() {
	mapTest()
}