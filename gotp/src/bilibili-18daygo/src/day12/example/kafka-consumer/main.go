package main

import (
	"fmt"
	"strings"
	"sync"
	"github.com/Shopify/sarama"
	"time"
)

var (
	wg = sync.WaitGroup{}
)

func main() {
	consumer, err := sarama.NewConsumer(strings.Split("192.168.31.177:9092", ","), nil)
	if err != nil {
		fmt.Println("failed to start consumer: ", err)
		return
	}
	partitionList, err := consumer.Partitions("nginx_log")
	if err != nil {
		fmt.Println("failed to get the list of partitions: ", err)
		return
	}
	fmt.Println(partitionList)

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("nginx_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d: %s\n",
				partition, err)
			return
		}
		defer pc.AsyncClose()
		go func(pc sarama.PartitionConsumer) {
			// wg.Add(1)
			for msg := range pc.Messages() {
				fmt.Printf("partition: %d, offset: %d, key: %s, value: %s\n",
					msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
			// wg.Done()
		}(pc)
	}
	
	// wg.Wait()
	time.Sleep(time.Hour)
	consumer.Close()
}
