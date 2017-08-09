/*
 * Revision History:
 *     Initial: 2017/06/20        Yusan Kurban
 */

package monitor

import (
	"fmt"
	"strings"

	"github.com/yusank/wechat-go/wxweb"
)

var (
	// Debug is channel for receiving message
	Debug chan int
	// TestText is for testing alive() if alive
	TestText = "areUalive"
)

func init() {
	Debug = make(chan int, 1000)
}

// Register register plugs
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
		fmt.Printf("Got a location message: %v", l[1])
		session.SendText("I`m location", session.Bot.UserName, wxweb.RealTargetUserName(session, msg))
	default:
		fmt.Printf("not a location message.")
	}
}

// Debugger listen Debug channel and when got any message from channel it will send to myself immediately
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
