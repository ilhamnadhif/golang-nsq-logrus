package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func(counter int) {
			fmt.Println(counter)
		}(i)
	}
	time.Sleep(10 * time.Second)
}
