package main

import(
	"encoding/json"
	"fmt"
	"os"
	"math/rand"
	"io/ioutil"
)


// Name string  `json:"name"`  第一个Name 要大写，name 要加引号""
type Student struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Score float64 `json:"score"`
}

func testMarshal(file string){
	// 准备好要序列化的interface{}对象
	var students []*Student
	for i:=0;i<10;i++ {
		name := fmt.Sprintf("name%d",i)
		age := rand.Intn(100)
		score := rand.Float64()*100

		Stu := Student{
			Name:name,
			Age:age,
			Score:score,
		}
		students = append(students, &Stu)
	}

	// 序列化得到byte切片
	dataBytes, err := json.Marshal(&students)
	if err != nil {
		fmt.Println("Marshal failed", err)
		return
	}

	// 将byte切片写到文件中
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file failed", err)
		return
	}
	defer f.Close()

	n, err := f.Write(dataBytes)
	if err != nil {
		fmt.Println("write file failed", err)
		return
	}

	fmt.Println("Write n:", n)

}

func testUmarshal(file string) {
	
	dataBytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("read file failed:", err)
		return
	}
	// 反序列化需要定义一个变量接收
	var students []*Student
	err = json.Unmarshal(dataBytes, &students)
	if err != nil {
		fmt.Println("json unmarshal failed:", err)
		return
	}

	fmt.Println(students)
	for i:=0;i<len(students);i++ {
		fmt.Println(students[i])
	}
}

func main() {
	file := "D:/students.json"
	// testMarshal(file)
	testUmarshal(file)
}