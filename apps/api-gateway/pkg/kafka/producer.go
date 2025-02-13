package kafka

import (
	"github.com/IBM/sarama"
	"log"
)

func ConnectProducer(config *sarama.Config) (sarama.SyncProducer, error) {
	// [todo] change the addrs
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Printf(err.Error())
	}
	return producer, nil
}

func ProduceMessage(topic string, message []byte) error {
	// [todo] change config
	producer, err := ConnectProducer(nil)
	if err != nil {
		log.Printf(err.Error())
	}
	defer producer.Close()
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message)}
	// Send the message to the topic
	_, _, err = producer.SendMessage(msg)
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	log.Printf("Message was sent successfully topic(%s) message (%s)", topic, message)
	return nil
}
