package main

import (
	"fmt"
	elastic "gopkg.in/olivere/elastic.v2"
)

type Nginx struct {
	User    string
	Message string
}

func main() {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://10.143.57.161:9200/"))
	if err != nil {
		fmt.Println("connect es error", err)
		return
	}
	fmt.Println("conn es succ")

	tweet := Nginx{
		User: "test", 
		Message: "i love go program, i am studying",
	}

	res, err := client.Index().
		Index("test").
		Type("tweet").   // 类型
		//Id("1").      // 好像可以随机设置
		BodyJson(tweet).
		Do()
	if err != nil {
		fmt.Printf("send to es error %d", err)
		return
	}

	fmt.Println("res:", res)
	fmt.Println("insert success")
}  