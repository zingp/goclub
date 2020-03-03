package main

import (
	"io"
	"net/http"
	"fmt"
	"text/template"
)

var temp *template.Template

type Person struct {
	Name string
	Age int
}

func Hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Hello world</h1>")
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("My-Header", "gogogo")

	p := &Person{
		Name:"冰冰",
		Age: 30,
	}
	err := temp.Execute(w, p)
	if err != nil {
		fmt.Println("execute template error:", err)
	}
}


func init(){
	t, err := template.ParseFiles("./index.html")
	if err != nil {
		fmt.Println("read template error:", err)
		return
	}
	temp = t
}
func main(){
	http.HandleFunc("/", Hello)
	http.HandleFunc("/index", Index)
	if err := http.ListenAndServe(":8888",nil); err != nil {
		fmt.Println("listen :8888 error", err)
	}
}