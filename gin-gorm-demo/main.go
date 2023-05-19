package main

import (
	"gin-gorm-demo/server"
	_ "net/http/pprof"
)

func main() {
	server.Run()
}
