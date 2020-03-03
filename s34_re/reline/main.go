package main

import (
	"strings"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	line := `2018/08/09 11:07:43 [E] [logs.go:90] [Sogou-Observer,data=1,system=2,Owner=OP][invalid URL escape "%\\\""][[]]`
	// reg := regexp.MustCompile(`[a-z]+`)
	// fmt.Printf("%q\n", reg.FindAllString(text, -1))
	reg := regexp.MustCompile(`2018/08/09 11:07(?P<part1>.+)\[Sogou-Observer,(?P<part2>.+),Owner=OP\]`)
	params := reg.FindStringSubmatch(line)
	if len(params) == 0 {
		return
	}
	itemStr := params[2]
	fmt.Println(itemStr)
	itemSlice := strings.Split(itemStr, ",")

	var kvMap = make(map[string]int)
	for _, item := range itemSlice {
		if len(item) == 0 {
			continue
		}
		kvSlice := strings.Split(item, "=")
		k := kvSlice[0]
		v, err := strconv.Atoi(kvSlice[1])
		if err != nil {
			continue
		}

		value, ok := kvMap[k]
		if !ok {
			kvMap[k] = v
			continue
		}
		kvMap[k] = value + v
	}
	fmt.Println(kvMap)
}