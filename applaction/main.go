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
	for {
		dur := 2
		time.Sleep(time.Duration(dur) * time.Minute)

		monitor.Debug <- dur
	}
}
