package main

import (
	"io"
	"net/http"
	"fmt"
	"log"
)

const form = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
    <form action="#" method="post" name="bar">
        <div>Name:<input type="text" name="user" /></div>
        <div>Passwd:<input type="text" name="password" /></div>
        <input type="submit" value="Login" />
    </form>
</body>
</html>`

func Hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Hello world</h1>")
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("My-Header", "gogogo")

	switch r.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		r.ParseForm()  // 必须要调用这个解析
		io.WriteString(w, r.Form["user"][0])
		io.WriteString(w, "\n")
		io.WriteString(w, r.FormValue("password"))
	}
}

func main(){
	http.HandleFunc("/", Hello)
	http.HandleFunc("/login", Login)
	if err := http.ListenAndServe(":8888",nil); err != nil {
		fmt.Println("listen :8888 error", err)
	}
}


// 捕获异常
func testPanic(){
	http.HandleFunc("/", LogPanics(Hello))
	http.HandleFunc("/login", LogPanics(Login))
	if err := http.ListenAndServe(":8888",nil); err != nil {
		fmt.Println("listen :8888 error", err)
	}
}

func LogPanics(handle http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		defer func(){
			if err := recover(); err != nil{
				log.Printf("[%v] caught panic:%v", r.RemoteAddr, err)
			}
		}()

		handle(w, r)
	}
}