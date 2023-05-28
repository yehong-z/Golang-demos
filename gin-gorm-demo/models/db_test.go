package models

import (
	"testing"

	"golang.org/x/sync/singleflight"
)

func init() {
	InitDB()
}

var group singleflight.Group

func op() (interface{}, error) {
	var total int64
	DB.Model(&UserLogin{}).Count(&total)
	return nil, nil
}

func BenchmarkParallelSingleFlightOP(b *testing.B) {
	b.SetParallelism(32)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			group.Do("key", op)
		}
	})
}

func BenchmarkParallelOP(b *testing.B) {
	b.SetParallelism(32)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			op()
		}
	})
}
