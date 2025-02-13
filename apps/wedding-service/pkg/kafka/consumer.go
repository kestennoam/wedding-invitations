package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"wedding-service/interfaces"
)

func ConnectConsumer(config *sarama.Config) (sarama.Consumer, error) {
	// [todo] change the addrs
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Printf(err.Error())
	}
	return consumer, nil
}

func ConsumeWeddingMessage(topic string) (interfaces.Wedding, error) {
	consumer, err := ConnectConsumer(nil)
	if err != nil {
		log.Printf(err.Error())
		return interfaces.Wedding{}, err
	}
	defer consumer.Close()
	partition, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Printf(err.Error())
	}
	log.Println("consuming topic")
	for {
		select {
		case msg := <-partition.Messages():
			var wedding interfaces.Wedding
			err := json.Unmarshal(msg.Value, &wedding)
			if err != nil {
				log.Println(err)
				continue
			}
			return wedding, nil
		case err := <-partition.Errors():
			log.Println(err)
			return interfaces.Wedding{}, err
		}
	}
}
