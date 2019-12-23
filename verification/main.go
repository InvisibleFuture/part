package verification

import (
	"math"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var (
	codes sync.Map
	queue chan code
)

type code struct {
	key  string
	time int64
}

func init() {
	queue = make(chan code, 1024)
	go Hourglass(queue)
}

func Hourglass(ch chan code) {
	var item code
	var t, next int64

	for {
		t = time.Now().Unix()
		item = <-ch
		if item.time > t {
			next = (item.time - t) + 2 // Avoid negative numbers
			time.Sleep(time.Duration(next) * time.Second)
		}
		codes.Delete(item.key)
	}
}

func Consume(key string) string {
	random := strconv.Itoa(int(math.Round(rand.Float64() * 99999)))
	codes.Store(key, random)
	queue <- code{key, time.Now().Unix()}
	return random
}

func Produce(key, value string) bool {
	if data, ok := codes.Load(key); !ok || data != value {
		return false
	}
	codes.Delete(key)
	return true
}
