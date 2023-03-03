package utils

import (
	"fmt"
	"testing"
)

func TestTimeFunc(t *testing.T) {
	/*
		"2019-01-19 15:04:05"
		"2019年12月19日 12:30:05"
		"2019.12.19 6:"
	*/
	time, err := ParserTimeStr("23.03.10 15:00")
	if err != nil {
		fmt.Println("Parse time error:", err)
		return
	}
	fmt.Println(FormatTime(time))
}
