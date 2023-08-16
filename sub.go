package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"
)

func main() {
	config := nsq.NewConfig()

	consumerProcess, err := nsq.NewConsumer("test", "process", config)
	if err != nil {
		logrus.Panic("Could not create consumer")
	}
	consumerLog, err := nsq.NewConsumer("test", "log", config)
	if err != nil {
		logrus.Panic("Could not create consumer")
	}

	consumerProcess.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		logrus.Println("NSQ process message received:")
		logrus.Println(string(message.Body))
		return nil
	}))
	consumerLog.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		logrus.Println("NSQ log message received:")
		logrus.Println(string(message.Body))
		return nil
	}))

	err = consumerProcess.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		logrus.Panic("Could not connect")
	}
	err = consumerLog.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		logrus.Panic("Could not connect")
	}

	// Menunggu sinyal untuk keluar
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Menghentikan konsumen
	consumerProcess.Stop()
	consumerLog.Stop()
}
