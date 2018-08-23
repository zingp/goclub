package main

import (
	"context"
	"time"
	"net/http"
	"fmt"
	"io/ioutil"
)
/*超时控制*/
type Result struct {
	r *http.Response
	err error
}

func process() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	tr := &http.Transport{}
	client := &http.Client{Transport:tr}
	c := make(chan Result, 1)
	req, err := http.NewRequest("GET","http://www.444.com", nil)
	if err != nil {
		fmt.Println("http request error:", err)
		return
	}
	// 主要控制该goroutine超时
	go func() {
		resp, err := client.Do(req)
		pack := Result{r:resp, err:err}
		c <- pack
	}()

	select {
	// 如果超时
	case <-ctx.Done():
		tr.CancelRequest(req)
		res :=<- c
		fmt.Println("timeOut! error:", res.err)
	case res:=<-c:
		defer res.r.Body.Close()
		out, _ := ioutil.ReadAll(res.r.Body)
		fmt.Printf("Server Response:%s", out)
	}
	return
}


func main(){
	process()
}