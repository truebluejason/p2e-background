package syncLog

import (
	"fmt"
	"sync"
)

func syncLog() func(string) {
	mutex := sync.Mutex{}
	return func(msg string) {
		mutex.Lock()
		fmt.Println(msg)
		mutex.Unlock()
	}
}

var Println func(string) = syncLog()