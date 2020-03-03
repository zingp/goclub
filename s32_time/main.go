package main

import (
	"fmt"
	"time"
)

func reInitTime(date string, hour string){
	// 20180807 00-23
	// 20180807 12
	// timeStamp = time.Now().Add(h).Format("2006-01-02_15")
	// 	timeDay = time.Now().Add(h).Format("20060102")
	// 	timeMon = time.Now().Add(h).Format("200601")
	timeStamp := fmt.Sprintf("%s-%s-%s_%s",  date[0:4], date[4:6],date[6:8], hour)
	timeDay :=  date
	timeMon := date[0:6]
	fmt.Println(timeStamp, timeDay, timeMon)
}

func main(){
	// start := time.Now().Unix()
	// time.Sleep(5 * time.Second)
	// now := time.Now().Unix()
	// fmt.Println(now -start)
	// fmt.Printf("start=%d  end=%d", start,now)
	// day:="20180807"
	// fmt.Println(day[0:4])

	tricker := time.NewTicker(time.Second)
	for range tricker.C {
		starTime := time.Now().Unix()
		startimeStr := time.Now().Format("2006/01/02 15:04:05")
		if starTime % 30  == 0 {			
			fmt.Println(startimeStr)
		}
	}

	fmt.Println(int64(300) % 10)
	
}