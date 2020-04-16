package main

import (
	"fmt"
	"sync"
)

var wordMap sync.Map

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func sortSyncMapByValue(m sync.Map) PairList {
	p := make(PairList, 0)
	m.Range(func(k, v interface{}) bool {
		key, ok1 := k.(string)
		value, ok2 := v.(int)
		if ok1 && ok2 {
			p = append(p, Pair{key, value})
		}
		return true
	})
	return p
}


func main() {
	// 写入
	wordMap.Store("表里", 1)
	wordMap.Store("学习", 2)
	// 读取
	if v, ok := wordMap.Load("表里"); ok {
		fmt.Println(v)
	}

	// 修改值
	key := "学习"
	v, ok := wordMap.Load(key)
	if ok {
		if value, ok := v.(int); ok {
			wordMap.Store(key, value+1)
		}
	} else {
		wordMap.Store(key, 1)
	}

	wordMap.Range(
		func(k, v interface{}) bool {
			fmt.Printf("%s %d\n", k, v)
			return true
		})
	p := sortSyncMapByValue(wordMap)
	for _, v := range p {
		fmt.Printf("%s : %d\n", v.Key, v.Value)
	}
}
