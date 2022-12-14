package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"
	"sync"
)

func main() {
	config := nsq.NewConfig()
	p, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		logrus.Fatal(err)
	}
	wg := new(sync.WaitGroup)
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(counter int) {
			err := p.Publish("bootcamp", []byte(fmt.Sprintf("pesan ke - %d", counter)))
			if err != nil {
				logrus.Error(err)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
