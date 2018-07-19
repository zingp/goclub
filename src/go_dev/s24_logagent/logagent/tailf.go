package main

import(
	"github.com/hpcloud/tail"
	"github.com/astaxie/beego/logs"
	"fmt"
	"sync"
	"strings"
)

type Tailer struct {
	tail *tail.Tail
	offset int64
	filename string
}

var tailMgr *TailMgr
var waitGroup sync.WaitGroup

type TailMgr struct {
	tailerMap map[string]*Tailer
	lock sync.Mutex
}

func NewTailMgr() (*TailMgr) {
	return &TailMgr{
		tailerMap: make(map[string]*Tailer, 16),
	}
}

func (t *TailMgr)AddLogFile(filename string)(err error){
	t.lock.Lock()
	defer t.lock.Unlock()

	_, ok := t.tailerMap[filename]
	if ok {
		err = fmt.Errorf("duplicate filename:%s", filename)
		return
	}

	tail, err := tail.TailFile(filename, tail.Config{
		ReOpen: true,
		Follow: true,
		Location:&tail.SeekInfo{ Offset:0, Whence: 2},  // 读到末尾
		MustExist: false,  //不存在也不报错
		Poll: true,
	})
	if err != nil {
		fmt.Println("tail file err:", err)
		return
	}

	tailer := &Tailer{
		filename:filename,
		offset:0,
		tail: tail,
	}

	t.tailerMap[filename] = tailer
	return
}

func (t *TailMgr) Process() {
	for _, tailObj := range t.tailerMap {
		waitGroup.Add(1)
		go tailObj.readLog()
	}
}

func (t *Tailer) readLog(){
	for line := range t.tail.Lines {
		if line.Err != nil {
			logs.Error("read line error:%v ", line.Err)
			continue
		}

		lineStr := strings.TrimSpace(line.Text)
		if len(lineStr)==0 ||lineStr[0]=='\n' {
			continue
		}
		
		kafkaSender.addMessage(line.Text)
	}
	waitGroup.Done()
}

func RunServer() {
	tailMgr = NewTailMgr()

	for _, filename := range appConfig.ListenFile {
		err := tailMgr.AddLogFile(filename)
		if err != nil {
			logs.Error("add logfile %v error:%v", filename, err)
			continue
		}
		logs.Debug("add logfile success %s", filename)
	}

	tailMgr.Process()
	waitGroup.Wait()
}