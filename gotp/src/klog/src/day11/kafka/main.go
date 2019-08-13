package main

import (
	"github.com/Shopify/sarama"
	"log"
)

func main() {
	config := sarama.NewConfig()

	client, err := sarama.NewSyncProducer([]string{"192.168.14.7:9092"}, config)
	if err != nil {
		log.Println("producer close, err: ", err)
		return
	}
	defer client.Close()

	msg := &sarama.ProducerMessage{}
	msg.Topic = "nginx_log"
	msg.Value = sarama.StringEncoder("this is a good test, my message is good")

	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		log.Println("send message fail, err: ", err)
		return
	}

	log.Printf("pid:%v, offset:%v\n", pid, offset)
}
