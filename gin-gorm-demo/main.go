package main

import (
	_ "net/http/pprof"

	"gin-gorm-demo/server"
)

func main() {
	server.Run()
}
