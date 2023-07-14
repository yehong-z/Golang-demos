package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

var DB *gorm.DB

func InitDB() {
	// 连接 MySQL 数据库
	dsn := "root:zyh130452@tcp(10.20.121.247:3306)/ranking_demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
}

type Video struct {
	gorm.Model
	VideoName  string
	ScoreCount int
	ScoreSum   int
	ScoreAvg   float64
}

func testAddCount(n int) {
	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				// 评分数目和评分的增量
				count := 1
				score := 5
				DB.Model(&Video{}).Where("id=?", "1").
					Updates(map[string]interface{}{
						"score_count": gorm.Expr("score_count + ?", count),
						"score_sum":   gorm.Expr("score_sum + ?", score),
						"score_avg":   gorm.Expr("(score_sum+?)/(score_count+?)", score, count),
					})
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func testBatchAdd() {
	count := 20
	score := 100

	for {
		select {
		case <-time.Tick(time.Second):
			go func() {
				start := time.Now()
				DB.Model(&Video{}).Where("id=?", "1").
					Updates(map[string]interface{}{
						"score_count": gorm.Expr("score_count + ?", count),
						"score_sum":   gorm.Expr("score_sum + ?", score),
						"score_avg":   gorm.Expr("(score_sum+?)/(score_count+?)", score, count),
					})
				duration := time.Since(start)
				fmt.Printf("Execution time: %v\n", duration)
			}()

		}
	}

}

func main() {
	InitDB()
	err := DB.AutoMigrate(&Video{})
	if err != nil {
		fmt.Println(err.Error())
	}
	testBatchAdd()
}
