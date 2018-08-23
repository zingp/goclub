package main

import(
	"net/http"
	"io/ioutil"
	"fmt"
)

func HttpGet(url string){
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(" http get error:", err)
		return
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("read date error:", err)
	}
	fmt.Println(string(data))
}
func main() {
	url := "https://www.cnblogs.com/zingp/p/6537841.html"
	HttpGet(url)

}