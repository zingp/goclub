package main


import(
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

type KafkaSender struct {
	client sarama.SyncProducer
	lineChan chan string
}

var kafkaSender =  &KafkaSender{}

func NewKafkaSender(kafkaAddr string)(kafka *KafkaSender, err error){
	kafka = &KafkaSender{
		lineChan: make(chan string, 100000),
	}

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll   //等待kafka ack
	config.Producer.Partitioner = sarama.NewRandomPartitioner  // 随机分区
	config.Producer.Return.Successes = true

	client, err := sarama.NewSyncProducer([]string{kafkaAddr}, config)
	if err!= nil {
		logs.Error("init kafka client err: %v", err)
		return
	}
	kafka.client = client

	for i:=0; i<appConfig.ThreadNum;i++ {
		go kafka.sendMsgToKfk()
	}
	return
}

func initKafka()(err error) {
	kafkaSender, err = NewKafkaSender(appConfig.KafkaAddr)
	return
}

func (k *KafkaSender) sendMsgToKfk() {

	for v := range k.lineChan {
		msg := &sarama.ProducerMessage{}
		msg.Topic = appConfig.KafkaTopic
		msg.Value = sarama.StringEncoder(v)

		_, _, err := k.client.SendMessage(msg)
		if err != nil {
			logs.Error("send massage to kafka error: %v", err)
			return
		}
	}
}

func (k *KafkaSender) addMessage(line string) (err error){
	k.lineChan <- line
	return
}


// func RunServer(){
// 	for i:=0;i<10;i++ {
// 		kafkaSender.addMessage(fmt.Sprintf("wo ai beijing tiananmen %d\n",i))
// 	}
// 	time.Sleep(10*time.Second)
// }