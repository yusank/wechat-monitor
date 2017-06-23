/*
 * Revision History:
 *     Initial: 2017/06/20        Yusan Kurban
 */

package monitor

import (
	"github.com/songtianyi/wechat-go/wxweb"
	"fmt"
)


var (
	Debug    chan int
)

func init() {
	Debug = make(chan int, 1000)
}


func Debugger(session *wxweb.Session, target string) {

	myself := session.Bot.UserName

	for {
		select {
		case d := <-Debug:
			s := fmt.Sprintf("Receive message %d", d)
			session.SendText(s, myself, target)
		}
	}
}
