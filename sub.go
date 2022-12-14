package main

import (
	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)

	decodeConfig := nsq.NewConfig()
	c, err := nsq.NewConsumer("bootcamp", "My_NSQ_Channel", decodeConfig)
	if err != nil {
		logrus.Panic("Could not create consumer")
	}

	c.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		logrus.Println("NSQ message received:")
		logrus.Println(string(message.Body))
		return nil
	}))
	err = c.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		logrus.Panic("Could not connect")
	}

	wg.Wait()
}
