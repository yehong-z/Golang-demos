package generator

import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

const (
	epoch    = int64(1577808000000) // 2020/1/1 00:00:00 UTC+8
	workerID = uint16(1)            // 机器 ID，可以根据需要修改
	maxSeq   = uint16(4096)         // 同一毫秒内最多允许 4096 个序号
)

var (
	lastTimestamp int64
	sequence      uint16
	mu            sync.Mutex
)

func getSnowFlakeID() (int64, error) {
	mu.Lock()
	defer mu.Unlock()
	timestamp := time.Now().UnixNano() / 1e6 // 毫秒级时间戳
	if timestamp == lastTimestamp {
		sequence++
		if sequence >= maxSeq {
			time.Sleep(time.Duration(rand.Intn(2)) * time.Millisecond)
			return getSnowFlakeID()
		}
	} else {
		sequence = 0
	}

	if timestamp < lastTimestamp {
		return 0, errors.New("Invalid timestamp")
	}

	lastTimestamp = timestamp

	id := ((timestamp - epoch) << 22) | (int64(workerID) << 12) | int64(sequence)
	return id, nil
}
