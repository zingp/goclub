package main
//冒泡排序，比Python快二十倍以上
import (
	"math/rand"
	"time"
	"fmt"
	//"sort"
)

func bubble(a []int) {
	for i:=0; i<(len(a)-1); i++ {
		status := false
		for j:=0; j<(len(a)-i-1); j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				status = true
			}
		}

		if !status {
			return
		}
	}
}

func main() {
	num := 10000
	var array []int
	for i:=0;i<num;i++{
		array = append(array, rand.Intn(10000))
	}
	fmt.Println("before:", array)

	start := time.Now().UnixNano()
	bubble(array[:])
	//sort.Ints(array[:])
	end := time.Now().UnixNano()
	fmt.Printf("cost:%d\n", end-start)

	fmt.Println("after:", array)

}
