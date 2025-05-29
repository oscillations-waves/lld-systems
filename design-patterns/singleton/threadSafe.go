package singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}
var threadSafeSingleInstance *threadSafeSingleton

type threadSafeSingleton struct {
}

func getThreadSafeInstance() *threadSafeSingleton {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating thread-safe single instance now.")
			threadSafeSingleInstance = &threadSafeSingleton{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}
	return threadSafeSingleInstance

}
