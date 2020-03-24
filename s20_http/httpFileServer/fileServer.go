package main

import(
	"net/http"
	// "fmt"
)

// 简单文件服务
func SimpleFileServer(){
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.ListenAndServe(":8081", nil)
}

// 
func MyFileServer(){
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.ListenAndServe(":8080", nil)
}

func main(){
	SimpleFileServer()
}