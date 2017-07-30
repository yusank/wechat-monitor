/*
 * Revision History:
 *     Initial: 2017/06/20        Yusan Kurban
 */

package monitor

import (
	"github.com/yusank/wechat-go/wxweb"
	"fmt"
	"strings"
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
	session.HandlerRegister.Add(wxweb.MSG_TEXT, wxweb.Handler(location), "loc")
}

func alive(session *wxweb.Session, msg *wxweb.ReceivedMessage) {
	switch msg.Content {
	case TestText:
		session.SendText("I`m alive", session.Bot.UserName, wxweb.RealTargetUserName(session, msg))

	default:
		fmt.Printf("Got a new message from: %s \n", session.Cm.GetContactByUserName(msg.FromUserName).PYQuanPin)
	}
}

func location(session *wxweb.Session, msg *wxweb.ReceivedMessage) {
	switch msg.SubType {
	case wxweb.MSG_LOCATION:
		l := strings.Split(msg.Url, "?")
		c := strings.Split(l[1], "")
		fmt.Printf("Got a location message: %v", l[1])
		session.SendText("I`m location", session.Bot.UserName, wxweb.RealTargetUserName(session, msg))
	default:
		fmt.Printf("Got a location message: %v AND %v \n", msg.Content, msg.OriginContent)
		session.SendText("I`m location", session.Bot.UserName, wxweb.RealTargetUserName(session, msg))
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
