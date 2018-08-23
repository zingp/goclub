package main

import(
	"io/ioutil"
	"fmt"
)

func main(){
	folder := "../.."
	listFile(folder)
}

func listFile(folder string){
	files, _ := ioutil.ReadDir(folder) //specify the current dir
	for _,file := range files{
			if file.IsDir(){
					listFile(folder + "/" + file.Name())
			}else{
					fmt.Println(folder + "/" + file.Name())
			}
	}
}