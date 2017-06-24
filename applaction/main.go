/*
 * Revision History:
 *     Initial: 2017/06/20        Yusan Kurban
 */

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

func Tester() {
	time.Sleep(60 * time.Second)
	for {
		dur := 2
		time.Sleep(time.Duration(dur) * time.Minute)

		monitor.Debug <- dur
	}
}
