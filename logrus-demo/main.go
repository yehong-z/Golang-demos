package main

import (
	"fmt"
	"logrus-demo/logger"
)

func init() {
}

type Log struct {
	source string
	data   any
}

type userInfo struct {
	userId   int
	username string
}
type and struct {
	a any
}

func main() {
	//logger := logger.NewLogger()
	//logger.Info("salkdfj")
	//
	//log := logrus.New()
	logger := logger.NewMiddlewareLogger()
	logger.Info(Log{
		source: "xxx",
		data: userInfo{
			username: "zyh",
			userId:   123,
		},
	})
	logger.Info(Log{
		source: "xxx",
		data: map[string]int{
			"aaa": 1,
			"bbb": 2,
		},
	})
	u := userInfo{
		username: "zyh",
		userId:   123,
	}
	fmt.Printf("%v", and{a: u})
	//server.Run()
}
