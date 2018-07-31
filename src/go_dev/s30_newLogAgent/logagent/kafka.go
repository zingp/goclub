package main

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)


type Message struct {
	line  string
	topic string
}

type KafkaSender struct {
	client   sarama.SyncProducer
	lineChan chan *Message
}

var kafkaSender = &KafkaSender{}

func NewKafkaSender(kafkaAddr string, threadNum int) (kafka *KafkaSender, err error) {
	kafka = &KafkaSender{
		lineChan: make(chan *Message, 10000),
	}

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          //等待kafka ack
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 随机分区
	config.Producer.Return.Successes = true

	client, err := sarama.NewSyncProducer([]string{kafkaAddr}, config)
	if err != nil {
		logs.Error("init kafka client err: %v", err)
		return
	}
	kafka.client = client

	for i := 0; i < threadNum; i++ {
		go kafka.sendMsgToKfk()
	}
	return
}

func initKafka(kafkaAddr string, threadNum int) (kafkaSender *KafkaSender, err error) {
	kafkaSender, err = NewKafkaSender(kafkaAddr, threadNum)
	return
}

func (k *KafkaSender) sendMsgToKfk() {

	for v := range k.lineChan {
		msg := &sarama.ProducerMessage{}
		msg.Topic = v.topic
		msg.Value = sarama.StringEncoder(v.line)

		_, _, err := k.client.SendMessage(msg)
		if err != nil {
			logs.Error("send massage to kafka error: %v", err)
			return
		}
	}
}

func (k *KafkaSender) addMessage(line string, topic string) (err error) {
	k.lineChan <- &Message{line: line, topic: topic}
	return
}
