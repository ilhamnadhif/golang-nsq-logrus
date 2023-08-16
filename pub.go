package main

import (
	"fmt"

	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"
)

func main() {
	config := nsq.NewConfig()
	p, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		logrus.Fatal(err)
	}
	for i := 1; i <= 10; i++ {
		err := p.Publish("test", []byte(fmt.Sprintf("pesan ke - %d", i)))
		if err != nil {
			logrus.Error(err)
		}
	}
	p.Stop()
}
