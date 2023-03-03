package reminder

import (
	"fmt"
	"time"

	"github.com/eatmoreapple/openwechat"
)

func NewReminder(self *openwechat.Self) {
	time.Sleep(time.Second * 5)
	entrys, err := loadRemindList()
	if err != nil {
		fmt.Println("Error: load remind list error: ", err)
		return
	}
	for _, entry := range entrys {
		sendTextToSelf(self, entry.content)
	}
}

func RemindMeRepayment(self *openwechat.Self) error {
	return sendTextToSelf(self, "发工资啦，还贷款啦！")
}
