package reminder

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"time"

	"github.com/eatmoreapple/openwechat"
	"github.com/yimadai/my_bot/utils"
)

var (
	SrcPath  = path.Join(os.Getenv("GOPATH"), "src/github.com/yimadai/my_bot")
	ConfPath = path.Join(SrcPath, "conf/remind_list")
)

type entry struct {
	time    time.Time
	content string
}

func loadRemindList() ([]*entry, error) {
	fmt.Println("ConfPath: ", ConfPath)

	fi, err := os.Open(ConfPath)
	if err != nil {
		return nil, fmt.Errorf("Error: %s\n", err)
	}
	defer fi.Close()

	entrys := make([]*entry, 0)
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		eventStr := string(a)
		fmt.Println("eventStr: ", eventStr)
		result := strings.Split(eventStr, "&")
		fmt.Println("result: ", result)
		fmt.Println("len(result): ", len(result))

		if len(result) != 2 {
			fmt.Println("Error: event format error: ", result)
			continue
		}
		eventTimeStr, content := strings.TrimSpace(result[0]), strings.TrimSpace(result[1])
		eventTime, err := utils.ParserTimeStr(eventTimeStr)
		if err != nil {
			fmt.Println("Error: event time format error: ", eventStr)
			continue
		}
		entrys = append(entrys, &entry{
			time:    eventTime,
			content: content,
		})
	}
	return entrys, nil
}

func sendTextToSelf(self *openwechat.Self, text string) error {
	err := self.SendTextToFriends(text, 0, &openwechat.Friend{
		User: self.User,
	})
	return err
}
