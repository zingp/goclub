package main

import (
	"fmt"
	"sync/atomic"
	"github.com/astaxie/beego/logs"
)

type Host struct {
	Domain string
	IP     string
}

type rsyncLog struct {
	host      *Host
	remoteLog string
	localLog  string
	rsyncCmd  string
}

type lzopLog struct {
	host     *Host
	localLog string
	lzoLog   string
	lzoCmd   string
}

type hadpLog struct {
	host   *Host
	lzoLog string
	hdfs   string
	putCmd string
}

func genRsyncLogObj(remoteAddr string, localAddr string, logName string) {

	// 获取一小时前日志的名称
	lastHourLog := fmt.Sprintf("%s.%s", appConf.logName, timeStamp)
	// rsync 命令格式
	rsyncCmdFmt := `rsync -avzP rsync.%s::odin%s%s %s`
	for k, v := range hostMap {
		rsyncCmd := fmt.Sprintf(rsyncCmdFmt, k, remoteAddr, lastHourLog, localAddr)

		hostObj := &Host{
			Domain: k,
			IP:     v,
		}
		rsyncLogObj := &rsyncLog{
			host:      hostObj,
			remoteLog: fmt.Sprintf(`%s%s`, remoteAddr, lastHourLog),
			localLog:  fmt.Sprintf(`%s%s`, localAddr, lastHourLog),
			rsyncCmd:  rsyncCmd,
		}

		rsyncChan <- rsyncLogObj
	}
}

func (r *rsyncLog) Process() {

	//fmt.Println("rsync:", r.rsyncCmd)
	_, err := ExecCmdLocal(r.rsyncCmd)
	if err != nil {
		logs.Error("execute cmd:%s error:%v", r.rsyncCmd, err)
		return
	}
	logs.Info("execute cmd:%s success", r.rsyncCmd)

	// statInfoLog.log.2018-07-31_14.resin48.shouji.zw.ted.10.142.71.193.statInfoLog.log.lzo
	lzoLogName := fmt.Sprintf(`%s.%s.%s.%s.%s.lzo`, appConf.logName, timeStamp, r.host.Domain, r.host.IP, appConf.logName)
	lzoLog := fmt.Sprintf(`%s%s`, appConf.lzopAddr, lzoLogName)
	lzoLogObj := &lzopLog{
		host:     r.host,
		localLog: r.localLog,
		lzoLog:   lzoLog,
		lzoCmd:   fmt.Sprintf(`lzop %s -o %s`, r.localLog, lzoLog),
	}

	lzopChan <- lzoLogObj
}

func (lzo *lzopLog) Process() {
	// fmt.Println("lzop:", lzo.lzoCmd)
	// _, err := ExecCmdLocal(lzo.lzoCmd)
	// if err != nil {
	// 	logs.Error("execute cmd:%s error:%v", lzo.lzoCmd, err)
	// 	return
	// }
	// logs.Info("execute cmd:%s success", lzo.lzoCmd)

	hdfsLogFile := fmt.Sprintf(`%s%s/%s/`, appConf.hdfs, timeMon, timeDay)
	hadpLogObj := &hadpLog{
		host:   lzo.host,
		lzoLog: lzo.lzoLog,
		hdfs:   hdfsLogFile,
		putCmd: fmt.Sprintf(`%s fs -put %s %s`, appConf.hadoopClient, lzo.lzoLog, hdfsLogFile),
	}

	hadpChan <- hadpLogObj
}

func (h *hadpLog) Process() {
	fmt.Println("put to hadoop:", h.putCmd)
	// _, err := ExecCmdLocal(h.putCmd)
	// if err != nil {
	// 	logs.Error("put to hadoop failed cmd:%s error:%v", h.putCmd, err)
	// 	return
	// }
	
	logs.Info("put to hadoop success cmd:%s", h.putCmd)
	atomic.AddInt32(&eixtFlagObj.Num, 1)
	
	if eixtFlagObj.Num == int32(len(hostMap)) {
		close(rsyncChan)
		close(lzopChan)
		close(hadpChan)
	}
}

func rsyncToLocal() {
	defer waitGroup.Done()
	
	for {
		select{
		case rLog, ok := <-rsyncChan:
			if !ok {
				return
			}
			rLog.Process()
		default:
		}
	}
}

func lzopToLocal() {
	defer waitGroup.Done()

	for {
		select {
		case lLog, ok :=<- lzopChan:
			if !ok {
				return
			}
			lLog.Process()
		default:
		}
	}
}

func putToHdfs() {
	defer waitGroup.Done()

	for {
		select {
		case hLog,ok :=<- hadpChan:
			if !ok {
				return
			}
			hLog.Process()
		default:
		}
	}
}