package monitor

import (
	"github.com/songtianyi/wechat-go/wxweb"
	"github.com/songtianyi/rrframework/logs"
)


func Monitor(session *wxweb.Session, s string) {
	logs.Info(s)

	session.SendText(s, session.Bot.UserName,
		"96d425393e0d5040b6d00846b099d1ad65fed9defd26d9aee2c217abc4534b61")
}



