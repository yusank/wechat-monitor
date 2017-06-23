/*
 * MIT License
 *
 * Copyright (c) 2017 SmartestEE Inc.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

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
	"fmt"
)

type WeChat struct {
	session *wxweb.Session
}

var (
	WXService *WeChat
	wxGroup = "ceshiweixin"
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
	time.Sleep(20 * time.Second)

	target := wx.session.Cm.GetContactByPYQuanPin(wxGroup)
	monitor.Debugger(wx.session, target.UserName)

	<- errChan
}
