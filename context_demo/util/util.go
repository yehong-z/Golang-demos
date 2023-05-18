package util

import (
	"os"
	"os/signal"
	"syscall"
)

func Block() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
	print("program exit")
}
