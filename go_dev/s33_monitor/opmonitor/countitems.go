package main

import (
	"bufio"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var countItemsMap = make(map[string]int)

func countItems(file string) (err error) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	part := getPart()
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read string err", err)
		}
		//得到一行
		reg := regexp.MustCompile(part)
		params := reg.FindStringSubmatch(line)
		if len(params) == 0 {
			continue
		}
		itemStr := params[2]
		itemSlice := strings.Split(itemStr, ",")
		for _, item := range itemSlice {
			if len(item) == 0 {
				continue
			}
			kvSlice := strings.Split(item, "=")
			k := strings.TrimSpace(kvSlice[0])
			v, err := strconv.Atoi(strings.TrimSpace(kvSlice[1]))
			if err != nil {
				continue
			}

			value, ok := countItemsMap[k]
			if !ok {
				countItemsMap[k] = v
				continue
			}
			countItemsMap[k] = value + v
		}
	}
	return
}

func getPart() (part string) {
	partStr := `%s(?P<part1>.+)\[Sogou-Observer,(?P<part2>.+),Owner=OP\]`
	m, _ := time.ParseDuration("-1m")
	timeMinStr := time.Now().Add(m).Format("2006/01/02 15:04")
	part = fmt.Sprintf(partStr, timeMinStr)
	return
}

func cron(file string) {
	tricker := time.NewTicker(time.Second)
	for range tricker.C {
		nowTime := time.Now().Unix()
		if nowTime % 60 < 1 {
			err := countItems(file)
			if err != nil {
				continue
			}

			logStr := "items ["
			for k, v := range countItemsMap {
				kvStr := fmt.Sprintf("%s=%d,", k, v)
				logStr += kvStr
			}
			logStr += "]"

			logs.Info(logStr)
			// 清空map
			countItemsMap = make(map[string]int)
		}
	}
	waitGroup.Done()
}
