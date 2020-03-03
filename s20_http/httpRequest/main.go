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
	url := "http://127.0.0.1:9090/items"
	HttpGet(url)

}