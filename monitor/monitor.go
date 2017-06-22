package monitor

import (
	"github.com/songtianyi/wechat-go/wxweb"
	"github.com/songtianyi/rrframework/logs"
	"fmt"
	"time"
)

func Register(session *wxweb.Session) {
	session.HandlerRegister.Add(wxweb.MSG_TEXT, wxweb.Handler(monitor), "monitor")
}

var id string

func monitor(session *wxweb.Session, msg *wxweb.ReceivedMessage) {
	logs.Info("receiveMsg >> ", msg.Content)
	logs.Info("from >> ", msg.FromUserName)
	logs.Info("to >> ", msg.ToUserName)

	sendMsg := "hello form monitor"
	defaultTarget := wxweb.RealTargetUserName(session, msg)
	myself := session.Bot.UserName
	//fileTarget := "filehelper"
	username := fmt.Sprintf("Message: %s::: From %s", msg.Content, session.Bot.NickName)
	friends := session.Cm.GetAll()
	friend := session.Cm.GetContactByPYQuanPin("Yusan")
	group := session.Cm.GetGroupContacts()

	switch msg.FromUserName {
	case myself:
		session.SendText(username, myself, friend.UserName)

		logs.Info("sendMsg >> ", username)
		logs.Info("from >> ", session.Bot.UserName)
		logs.Info("to >> ", session.Bot.NickName)


	default:
		for {
			time.Sleep(5 * time.Second)

			session.SendText(sendMsg, myself, defaultTarget)

			logs.Info("sendMsg >> ", sendMsg)
			logs.Info("from >> ", session.Bot.UserName, msg.RecommendInfo)
			logs.Info("to >> ", defaultTarget)

			for i, f := range friends {
				logs.Info("friends no. >> ", i, "nickname:", f.NickName,"disname:",
					f.DisplayName, "pinyin:", f.PYQuanPin, "rmark:", f.RemarkPYQuanPin)
			}
			for j, g := range group {
				logs.Info("friends no. >> ", j, "nickname:", g.NickName,"disname:",
					g.DisplayName, "pinyin:", g.PYQuanPin, "rmark:", g.RemarkPYQuanPin)
			}

		}
	}




}



