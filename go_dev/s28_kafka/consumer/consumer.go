package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"strings"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func main() {
	consumer, err := sarama.NewConsumer(strings.Split("10.134.123.183:9092", ","), nil)
	if err != nil {
		fmt.Printf("Failed to start consumer: %s", err)
		return
	}

	partitionList, err := consumer.Partitions("nginx_log")
	if err != nil {
		fmt.Println("Failed to get the list of partitions: ", err)
		return
	}
	fmt.Println(partitionList)

	for partition := range partitionList {
		fmt.Println("hello")
		pc, err := consumer.ConsumePartition("nginx_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("Failed to start consumer for partition %d: %s\n", partition, err)
			return
		}
		defer pc.AsyncClose()

		wg.Add(1)
		go func(p sarama.PartitionConsumer) {
			for msg := range p.Messages() {
				// p.Messages() 这是一个管道会阻塞
				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				fmt.Println()
				time.Sleep(time.Millisecond)
			}

			wg.Done()
		}(pc)
	}

	wg.Wait()
	consumer.Close()
}
