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
			lineSlice := strings.Split(line, ":")
			hostMap[lineSlice[0]] = lineSlice[1]
			break
		}
		if err != nil {
			continue
		}
		line = strings.TrimSpace(line)
		lineSlice := strings.Split(line, ":")
		hostMap[lineSlice[0]] = lineSlice[1]
	}
	return
}
