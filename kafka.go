package mjhp

import (
	"github.com/optiopay/kafka"
	"encoding/json"
	"log"
	"github.com/optiopay/kafka/proto"
)

var sendChan chan *JudgeBatchResult
var sendRun = true

var KEY_MJHP_RESP = []byte("mjhp.resp")

func ShutdownSendWorker() {
	sendRun = false
}

func StartKafka() {
	brokerConf := kafka.NewBrokerConf("mjhp-client")
	brokerConf.AllowTopicCreation = true
	broker, err := kafka.Dial(cfg.KafkaAddr, brokerConf)
	if err != nil {
		panic(err)
	}
	defer broker.Close()
	log.Println("subscribe topic: ", cfg.KafkaTopic)
	conf := kafka.NewConsumerConf(cfg.KafkaTopic, 0)
	conf.MinFetchSize = 1
	conf.StartOffset = kafka.StartOffsetNewest
	log.Println("Consumer conf: ", conf)
	consumer, err := broker.Consumer(conf)
	if err != nil {
		panic(err)
	}
	go startProducerWork(broker)
	for {
		msg, err := consumer.Consume()
		if err != nil {
			if err == kafka.ErrNoData {
				continue
			}
			panic(err)
		}
		judge := &JudgeReqBatch{}
		err = json.Unmarshal(msg.Value, judge)
		if err != nil {
			log.Println("recv unknown msg: ", msg.Value)
		} else {
			workChan <- judge
		}
	}
}
func startProducerWork(broker *kafka.Broker) {
	conf := kafka.NewProducerConf()
	conf.RequiredAcks = proto.RequiredAcksLocal
	producer := broker.Producer(conf)
	sendChan = make(chan *JudgeBatchResult, 100)
	for sendRun {
		select {
		case resp := <-sendChan:
			bin, err := json.Marshal(resp)
			if err != nil {
				log.Println("json marshal error: ", err)
				continue
			}
			msg := &proto.Message{Key: KEY_MJHP_RESP, Value: bin}
			log.Printf("resp to: %s, data: %v", resp.fromTopic, resp)
			_, err = producer.Produce(resp.fromTopic, 0, msg)
			if err != nil {
				log.Println("kafka produce error: ", err)
			}
		}
	}
}
