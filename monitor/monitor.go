package monitor

import (
	"github.com/songtianyi/wechat-go/wxweb"
	"github.com/songtianyi/rrframework/logs"
	"fmt"
)

func Register(session *wxweb.Session) {
	session.HandlerRegister.Add(wxweb.MSG_TEXT, wxweb.Handler(monitor), "monitor")
}

var (
	secertID = "yusan112233"
	Debug    chan int
)

func init() {
	Debug = make(chan int, 1000)
}

func monitor(session *wxweb.Session, msg *wxweb.ReceivedMessage) {
	defaultTarget := wxweb.RealTargetUserName(session, msg)

	switch msg.Content {
	case secertID:
		debugger(session, defaultTarget)
	default:
		s := fmt.Sprintf("receive message>>:%s from:%s", msg.Content, msg.FromUserName)
		logs.Info(s)
	}
}

func debugger(session *wxweb.Session, target string) {

	myself := session.Bot.UserName

	for {
		select {
		case d := <-Debug:
			s := fmt.Sprintf("Receive message %d", d)
			session.SendText(s, myself, target)
		}
	}
}
