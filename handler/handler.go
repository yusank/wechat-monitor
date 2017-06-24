/*
 * Revision History:
 *     Initial: 2017/06/20        Yusan Kurban
 */

package handler

import (
	"wechat-monitor/monitor"
	"github.com/songtianyi/wechat-go/wxweb"
	"github.com/songtianyi/rrframework/logs"

	"time"
)

type WeChat struct {
	session *wxweb.Session
}

var (
	WXService *WeChat
	WXGroup = "ceshiweixin"
)

func init() {
	WXService = &WeChat{
		session: &wxweb.Session{},
	}
}

func (wx *WeChat) HandleDebug() {
	var err error

	wx.session, err = wxweb.CreateSession(nil, nil, wxweb.TERMINAL_MODE)
	if err != nil {
		logs.Error(err)
		return
	}

	monitor.Register(wx.session)
	wx.session.HandlerRegister.EnableByName("testAlive")

	errChan := make(chan error)
	go func(errChan chan error) {
		if err := wx.session.LoginAndServe(false); err != nil {
			logs.Error("session exit, %s", err)
			errChan <- err
		}
	}(errChan)

	tick := time.Tick(2 * time.Second)
	for {
		select {
		case err := <-errChan:
			panic(err)
		case <-tick:
			// check login status
			if wx.session.Cookies != nil {
				break
			}
		}
		break

	}
	// now it is ok to send message
	// sleep 20 second for waiting login complete
	time.Sleep(60 * time.Second)

	target := wx.session.Cm.GetContactByPYQuanPin(WXGroup)
	monitor.Debugger(wx.session, target.UserName)

	<- errChan
}
