package main

import (
	"encoding/json"
	"fmt"
)

/*
有时候结构体中为了数据安全都用小写。怎么序列化？
用tag
*/

// 不用tag只能识别大写。
type stu struct {
	Name string
	Age int
	Score int
}
// 好像必须大写才行，json的时候吧大写转换为小写。
type stuTag struct {
	Name string `json:"name"`
	Age int     `json:"age"`
	Score int   `json:"score"`
}

func main() {
	stu01 := stu{
		Name:"zing-p",
		Age:23,
		Score:80,
	}
	stu02 := stuTag {
		Name:"liuYY",
		Age:25,
		Score:90,
	}

	data01, err := json.Marshal(stu01)
	if err != nil {
		fmt.Println("json encode stu01 failed:",err)
	}
	fmt.Printf(" json stu01:%s\n",data01)

	data02, err := json.Marshal(stu02)
	if err != nil {
		fmt.Println("json encode stu02 failed:",err)
	}
	fmt.Printf(" json stu02:%s\n",data02)

}
/*
运行结果：
json stu01:{"Name":"zing-p","Age":23,"Score":80}
json stu02:{"name":"liuYY","age":25,"score":90}
*/
