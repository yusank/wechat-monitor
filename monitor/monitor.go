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
	TestText = "areUalive"
)

func init() {
	Debug = make(chan int, 1000)
}

func Register(session *wxweb.Session) {
	session.HandlerRegister.Add(wxweb.MSG_TEXT, wxweb.Handler(alive), "testAlive")
}

func alive(sesion *wxweb.Session, msg *wxweb.ReceivedMessage) {
	switch msg.Content {
	case TestText:
		sesion.SendText("I`m alive", sesion.Bot.UserName, wxweb.RealTargetUserName(sesion, msg))

	default:
		fmt.Printf("Got a new message from: %s \n", sesion.Cm.GetContactByUserName(msg.FromUserName).PYQuanPin)
	}
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
