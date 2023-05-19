package server

import (
	"fmt"
	"gin-gorm-demo/config"
	"gin-gorm-demo/router"
	"log"
	"net/http"
	"os"
	"runtime"
)

func Run() {
	runtime.GOMAXPROCS(1)              // 限制 CPU 使用数，避免过载
	runtime.SetMutexProfileFraction(1) // 开启对锁调用的跟踪
	runtime.SetBlockProfileRate(1)     // 开启对阻塞操作的跟踪
	go func() {
		// 启动一个 http server，注意 pprof 相关的 handler 已经自动注册过了
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	r := router.InitRouter()
	err := r.Run(fmt.Sprintf("%v:%v", config.Info.IP, config.Info.Port))
	if err != nil {
		return
	}
}
