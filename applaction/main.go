package main

import (
	"wechat-monitor/handler"
	"wechat-monitor/monitor"
	"time"
)


func main() {
	go Tester()
	handler.WXService.HandleDebug()
}

func Tester(){
	for i := 0; i < 11; i++ {
		time.Sleep(5 * time.Second)

		monitor.Debug <- i
	}
}