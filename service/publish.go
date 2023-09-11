package service

import (
	"log"

	"github.com/nats-io/nats.go"
)

func PublishMsg(subject, msg string) string {
	nc, err := nats.Connect("nats://nats:4222")
	if err != nil {
		log.Fatal(err)
		return "Fail"
	}

	defer nc.Close()

	if err := nc.Publish("notebook", []byte("create")); err != nil {
		log.Fatal(err)
		return "Fail"
	}
	return "Ok"
}

// func main() {

// 	url := os.Getenv("NATS_URL")
// 	if url == "" {
// 		url = nats.DefaultURL
// 	}

// 	nc, _ := nats.Connect(url)

// 	defer nc.Drain()

// 	nc.Publish("greet.joe", []byte("hello"))

// 	sub, _ := nc.SubscribeSync("greet.*")

// 	msg, _ := sub.NextMsg(10 * time.Millisecond)
// 	fmt.Println("subscribed after a publish...")
// 	fmt.Printf("msg is nil? %v\n", msg == nil)

// 	nc.Publish("greet.joe", []byte("hello"))
// 	nc.Publish("greet.pam", []byte("hello"))

// 	msg, _ = sub.NextMsg(10 * time.Millisecond)
// 	fmt.Printf("msg data: %q on subject %q\n", string(msg.Data), msg.Subject)

// 	msg, _ = sub.NextMsg(10 * time.Millisecond)
// 	fmt.Printf("msg data: %q on subject %q\n", string(msg.Data), msg.Subject)

// 	nc.Publish("greet.bob", []byte("hello"))

// 	msg, _ = sub.NextMsg(10 * time.Millisecond)
// 	fmt.Printf("msg data: %q on subject %q\n", string(msg.Data), msg.Subject)
// }
