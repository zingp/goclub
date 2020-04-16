package main

import (
	"sort"
	"sync"
)

/* Map 排序 */
//要对golang map按照value进行排序，思路是直接不用map。
//用struct存放key和value，实现sort接口，就可以调用sort.Sort进行排序。
// A data structure to hold a key/value pair.
type Pair struct {
    Key   string
    Value int
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p PairList) Len() int           { return len(p) }

func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

// A function to turn a map into a PairList, then sort and return it.
func sortMapByValue(m map[string]int) PairList {
    p := make(PairList, len(m))
    i := 0
    for k, v := range m {
		p[i] = Pair{k, v}
		i += 1
	}
	// 降序
	sort.Sort(sort.Reverse(p))
	// sort.Sort(p)  // 升序
    return p
}

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
	// 降序
	sort.Sort(sort.Reverse(p))
	// sort.Sort(p)  // 升序
    return p
}
