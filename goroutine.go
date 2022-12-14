package main

import (
	"fmt"
	"time"
)

func Display(i int) {
	fmt.Println(i)
}

func main() {
	for i := 0; i < 10; i++ {
		go Display(i)
	}
	time.Sleep(10 * time.Second)
}
