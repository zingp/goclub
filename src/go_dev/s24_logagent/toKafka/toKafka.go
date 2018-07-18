package main

import(
	"github.com/Shopify/sarama"
	"fmt"
	"time"
)

func testToKafka(){
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll   //等待kafka ack
	config.Producer.Partitioner = sarama.NewRandomPartitioner  // 随机分区
	config.Producer.Return.Successes = true

	msg := &sarama.ProducerMessage{}
	msg.Topic = "nginx_log"
	msg.Value = sarama.StringEncoder("this is test log line, very good.")

	client, err := sarama.NewSyncProducer([]string{"10.134.123.183:9092"}, config)
	if err!= nil {
		fmt.Println("producer close, err:", err)
		return
	}

	defer client.Close()

	for {
		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("send massage error:", err)
			return
		}
 
		fmt.Printf("pid=%v  offset=%v\n", pid, offset)
		time.Sleep(5 *time.Millisecond)
	}
}

func main(){
	testToKafka()
}