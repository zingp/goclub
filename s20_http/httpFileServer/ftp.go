package main

import (
	"io"
	"os"
	"net/http"
	"log"
)

//假设文件上传为本地的服务器，上传的基础路径
const BaseUploadPath = "./"


func handleUpload (w http.ResponseWriter, request *http.Request)  {
	//文件上传只允许POST方法
	if request.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("Method not allowed"))
		return
	}

	//从表单中读取文件
	file, fileHeader, err := request.FormFile("file")
	if err != nil {
		_, _ = io.WriteString(w, "Read file error")
		return
	}
	//defer 结束时关闭文件
	defer file.Close()
	log.Println("filename: " + fileHeader.Filename)

	//创建文件
	newFile, err := os.Create(BaseUploadPath + "/" + fileHeader.Filename)
	if err != nil {
		_, _ = io.WriteString(w, "Create file error")
		return
	}
	//defer 结束时关闭文件
	defer newFile.Close()

	//将文件写到本地
	_, err = io.Copy(newFile, file)
	if err != nil {
		_, _ = io.WriteString(w, "Write file error")
		return
	}
	_,_ = io.WriteString(w, "Upload success")
}

// 文件下载处理器
func handleDownload (w http.ResponseWriter, request *http.Request) {
	// 文件下载只允许GET方法
	if request.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("Method not allowed"))
		return
	}
	// 文件名
	filename := request.FormValue("filename")
	if filename == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = io.WriteString(w, "Bad request")
		return
	}
	log.Println("filename: " + filename)
	//打开文件
	file, err := os.Open(BaseUploadPath + "/" + filename)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = io.WriteString(w, "Bad request")
		return
	}
	//结束后关闭文件
	defer file.Close()

	//设置响应的header头
	w.Header().Add("Content-type", "application/octet-stream")
	w.Header().Add("content-disposition", "attachment; filename=\""+filename+"\"")
	//将文件写至responseBody
	_, err = io.Copy(w, file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = io.WriteString(w, "Bad request")
		return
	}
}


// func main() {
// 	http.HandleFunc("/upload", handleUpload)
// 	http.HandleFunc("/download", handleDownload)

//     log.Println("Start listening :3000")
// 	err := http.ListenAndServe("0.0.0.0:3000", nil)
// 	if err != nil  {
// 		log.Fatal("Server run fail")
// 		return
// 	} 
// }


