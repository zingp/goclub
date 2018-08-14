package main

import (
	"bufio"
	"github.com/astaxie/beego/logs"
	"io"
	"os"
	"strings"
)

// GetHostMap func get host return map
func GetHostMap(file string) (err error) {
	f, err := os.Open(file)
	if err != nil {
		logs.Error("open file:%s error:%v", file, err)
		return
	}
	defer f.Close()

	hostMap = make(map[string]string, 60)
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			line = strings.TrimSpace(line)
			if len(line) == 0 {
				break
			}
			lineSlice := strings.Split(line, ":")
			if len(lineSlice) == 2 {
				hostMap[lineSlice[0]] = lineSlice[1]
				break
			}
			logs.Error("host conf error:%v", lineSlice)
		}
		if err != nil {
			continue
		}
		
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		lineSlice := strings.Split(line, ":")
		if len(lineSlice) == 2 {
			hostMap[lineSlice[0]] = lineSlice[1]
			continue
		}
		logs.Error("host conf error:%v", lineSlice)
	}
	return
}
