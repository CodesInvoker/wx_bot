package main

import (
	"fmt"

	"github.com/eatmoreapple/openwechat"
	"github.com/yimadai/my_bot/reminder"
)

func main() {
	defer func() {
		var err error
		if p := recover(); p != nil {
			switch p := p.(type) {
			case error:
				err = p
			default:
				err = fmt.Errorf("%s", p)
			}
		}
		fmt.Println("error at main", err)
	}()
	// bot := openwechat.DefaultBot()
	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式，上面登录不上的可以尝试切换这种模式

	// 注册消息处理函数
	bot.MessageHandler = func(msg *openwechat.Message) {
		if msg.IsText() && msg.Content == "ping" {
			msg.ReplyText("pong")
		}
	}
	// 注册登陆二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	// 登陆方式一：扫码登陆
	// if err := bot.Login(); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// 登录方式二：热登录，创建热存储容器对象。减少扫码次数
	reloadStorage := openwechat.NewJsonFileHotReloadStorage("storage.json")
	defer reloadStorage.Close()
	// 执行热登录
	if err := bot.HotLogin(reloadStorage, openwechat.NewRetryLoginOption()); err != nil {
		fmt.Println(err)
		return
	}

	// 获取登陆的用户
	self, err := bot.GetCurrentUser()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 获取所有的好友
	friends, err := self.Friends()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(friends)

	// 获取所有的群组
	groups, err := self.Groups()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(groups)

	go func() {
		// err = reminder.RemindMeRepayment(self)
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		reminder.NewReminder(self)
	}()

	// 阻塞主goroutine, 直到发生异常或者用户主动退出
	bot.Block()
}
