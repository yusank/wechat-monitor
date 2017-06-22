package main

import (
	"wechat-monitor/handler"
	//"time"
)


func main() {
	handler.WXService.HandleDebug()

	//for i := 0; i < 10; i++ {
	//	time.Sleep(3 * time.Second)
	//	if i % 2 != 0 {
	//		continue
	//	}
	//	handler.WXService.SendText("hello from go.")
	//
	//}
}
