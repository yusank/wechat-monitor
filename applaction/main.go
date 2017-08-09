/*
 * Revision History:
 *     Initial: 2017/06/20        Yusan Kurban
 */

package main

import (
	"time"

	"wechat-monitor/handler"
	"wechat-monitor/monitor"
)

func main() {
	go Tester()
	handler.WXService.HandleDebug()
}

// Tester sleep 1 minute for waiting login complete and send 2 to Debug channel every 2 second
func Tester() {
	time.Sleep(60 * time.Second)
	for {
		dur := 2
		time.Sleep(time.Duration(dur) * time.Minute)

		monitor.Debug <- dur
	}
}
