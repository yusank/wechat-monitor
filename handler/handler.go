/*
 * Revision History:
 *     Initial: 2017/06/20        Yusan Kurban
 */

package handler

import (
	"time"

	"github.com/songtianyi/rrframework/logs"
	"github.com/yusank/wechat-go/wxweb"

	"wechat-monitor/monitor"
)

// WeChat wrapper session
type WeChat struct {
	session *wxweb.Session
}

var (
	// WXService is provide interface for WeChat`s function
	WXService *WeChat
	// WXGroup is the wechat group name where message to be send
	WXGroup = "ceshiweixin"
)

func init() {
	WXService = &WeChat{
		session: &wxweb.Session{},
	}
}

// HandleDebug handle wechat login and refresh
func (wx *WeChat) HandleDebug() {
	var err error

	wx.session, err = wxweb.CreateSession(nil, nil, wxweb.TERMINAL_MODE)
	if err != nil {
		logs.Error(err)
		return
	}

	monitor.Register(wx.session)
	wx.session.HandlerRegister.EnableByName("testAlive")
	wx.session.HandlerRegister.EnableByName("loc")

	errChan := make(chan error)
	go func(errChan chan error) {
		if err := wx.session.LoginAndServe(false); err != nil {
			errChan <- err
		}
	}(errChan)

	tick := time.Tick(2 * time.Second)
	for {
		select {
		case err := <-errChan:
			logs.Error("session exit:", err)
			for i := 0; i < 3; i++ {
				if err = wx.session.LoginAndServe(true); err != nil {
					logs.Error("re-login error:", err)
				}
				time.Sleep(3 * time.Second)
			}
			if wx.session, err = wxweb.CreateSession(nil, wx.session.HandlerRegister, wxweb.TERMINAL_MODE); err != nil {
				logs.Error("create new session is failed:", err)
			}
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

	<-errChan
}
