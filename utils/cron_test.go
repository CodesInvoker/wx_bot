package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"testing"
	"time"

	"github.com/robfig/cron/v3"
)

/*
	文档：https://pkg.go.dev/github.com/robfig/cron/v3
*/
func TestCronFunc(t *testing.T) {
	c := cron.New()
	fmt.Println("location:", c.Location())
	c.AddFunc("30 * * * *", func() { fmt.Println("Every hour on the half hour") })
	id, err := c.AddFunc("10 57 17 2 3 *", func() { fmt.Println("3月2号17点57分10秒") })
	if err != nil {
		fmt.Println("error:", err)
	}
	c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") })
	c.AddFunc("@every 10s", func() { fmt.Println("Every 10 sec") })

	c.Start()
	fmt.Println(fmt.Printf("cron entry:%+v", c.Entry(id)))
	fmt.Println(fmt.Printf("cron entry:%+v", c.Entries()))

	time.Sleep(time.Second * 1000)
}

func TestCronWithSecondFunc(t *testing.T) {
	c := cron.New(cron.WithSeconds())
	id, err := c.AddFunc("10 57 17 2 3 *", func() { fmt.Println("3月2号17点57分10秒") })
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(fmt.Printf("cron entry:%+v", c.Entry(id)))
	c.Start()
	time.Sleep(time.Second * 1000)
}

var (
	SrcPath  = path.Join(os.Getenv("GOPATH"), "src/github.com/yimadai/my_bot")
	ConfPath = path.Join(SrcPath, "conf/remind_list")
)

func TestRemindList(t *testing.T) {
	fmt.Println("ConfPath: ", ConfPath)

	fi, err := os.Open(ConfPath)
	if err != nil {
		fmt.Println(fmt.Printf("Error: %s\n", err))
		return
	}
	defer fi.Close()
	eventTimeStrList := make([]string, 0)
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
		eventTimeStrList = append(eventTimeStrList, strings.TrimSpace(result[0]))
	}
	c := cron.New()
	ids := make([]cron.EntryID, len(eventTimeStrList))
	for i, timeStr := range eventTimeStrList {
		id, err := c.AddFunc(timeStr, func() { fmt.Println(timeStr) })
		if err != nil {
			fmt.Println("error:", err)
		}
		ids[i] = id
	}
	c.Start()
	for i, id := range ids {
		fmt.Println(fmt.Printf("cron:%s next_time:%+v", eventTimeStrList[i], c.Entry(id)))
		c.Remove(id)
	}
	time.Sleep(time.Second * 1000)
}
