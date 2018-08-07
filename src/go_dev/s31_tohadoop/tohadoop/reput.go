package main

import (
	"fmt"
)

func reInitTime(date string, hour string){
	// date: 20060102
	// hour: 00 -23
	timeStamp = fmt.Sprintf("%s-%s-%s_%s",  date[0:4], date[4:6],date[6:8], hour)
	timeDay =  date
	timeMon = date[0:6]
}
