package mylog

import (
	"fmt"
)

func Message(messages ...interface{}) {
	for _, message := range messages {
		fmt.Printf("\033[1;31;40m%v\033[0m\n", message)
	}
}
