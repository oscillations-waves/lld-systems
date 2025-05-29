package singleton

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 10; i++ {
		go getInstance()
		time.Sleep(100 * time.Millisecond)
		go getThreadSafeInstance()
	}
	fmt.Scanln()
}
