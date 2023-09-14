package service

import (
	"log"

	"github.com/nats-io/nats.go"
)

func PublishMsg(subject, action string) string {
	nc, err := nats.Connect("nats://nats:4222")
	if err != nil {
		log.Fatal(err)
		return "Fail"
	}

	defer nc.Close()

	if err := nc.Publish(subject, []byte(action)); err != nil {
		log.Fatal(err)
		return "Fail"
	}
	return "Ok"
}
