package kafka

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"log"
	"wedding-service/interfaces"
)

func ConnectConsumer(config *sarama.Config) (sarama.Consumer, error) {
	if config == nil {
		config = sarama.NewConfig()
	}
	// TODO extract to env
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Printf("Error creating consumer: %v", err)
		return nil, err
	}
	return consumer, nil
}

func ConsumeWeddingMessage(topic string) (interfaces.Wedding, error) {
	consumer, err := ConnectConsumer(nil)
	if err != nil {
		log.Printf("Error connecting consumer: %v", err)
		return interfaces.Wedding{}, err
	}
	defer consumer.Close()

	partition, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Printf("Error consuming partition: %v", err)
		return interfaces.Wedding{}, err
	}
	defer partition.Close()

	log.Println("Consuming topic:", topic)
	for {
		select {
		case msg := <-partition.Messages():
			var wedding interfaces.Wedding
			err := json.Unmarshal(msg.Value, &wedding)
			if err != nil {
				log.Printf("Error unmarshalling message: %v", err)
				continue
			}
			return wedding, nil
		case err := <-partition.Errors():
			log.Printf("Error in partition: %v", err)
			return interfaces.Wedding{}, err
		}
	}
}
