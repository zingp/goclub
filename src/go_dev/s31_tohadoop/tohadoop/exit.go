package main

type eixtFlag struct {
	Num int32
	Flag map[string]chan bool
}
var rsyncExitChan = make(chan bool, 1)
var lzopExitChan = make(chan bool, 1)
var hadpExitChan = make(chan bool, 1)

var eixtFlagObj = &eixtFlag {
	Flag:make(map[string]chan bool, 3),
}