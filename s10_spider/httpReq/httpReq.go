package httpReq

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func GetHtml(url string) {
	// get 请求
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))

}
