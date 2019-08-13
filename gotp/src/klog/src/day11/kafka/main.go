package main

import (
	"github.com/Shopify/sarama"
	"log"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
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
