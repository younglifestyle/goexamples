package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"math/rand"
	"strconv"
	"time"
)

var (
	brokerList     = []string{"172.16.9.229:9029"}
	consumerConfig = sarama.NewConfig()
	producerConfig = sarama.NewConfig()
	partitions     = []int32{0, 1, 2, 3, 4, 5}
)

func startProducers() {

	client, clientErr := sarama.NewClient(brokerList, producerConfig)

	if clientErr != nil {
		fmt.Println("Error:", clientErr.Error())
		return
	}

	defer client.Close()

	newProducer, producerError := sarama.NewAsyncProducerFromClient(client)

	if producerError != nil {
		fmt.Println("Error:", producerError.Error())
		return
	}

	defer newProducer.Close()

	inputChannel := newProducer.Input()
	errorsChannel := newProducer.Errors()
	successChannel := newProducer.Successes()

	for {
		select {
		case successMessage := <-successChannel:
			fmt.Println("Message Published in Partition:", successMessage.Partition)
		case newErr := <-errorsChannel:
			fmt.Println(newErr.Error())
		case <-time.After(500 * time.Millisecond):

			data := sarama.StringEncoder(strconv.Itoa(int(time.Now().Unix())))

			partitionKey := partitions[rand.Intn(len(partitions))]

			newMessage := &sarama.ProducerMessage{
				Topic:     "users.recording",
				Partition: partitionKey,
				Key:       data,
				Value:     data,
			}

			fmt.Println("Publishing Message in Partition:", newMessage.Partition)

			inputChannel <- newMessage
		}
	}
}

func initProducerConfig() {
	producerConfig.Producer.Partitioner = sarama.NewManualPartitioner
	producerConfig.Producer.Compression = sarama.CompressionLZ4
	producerConfig.Producer.RequiredAcks = sarama.WaitForLocal
	producerConfig.Producer.Flush.Frequency = 1 * time.Second
	producerConfig.Producer.Retry.Max = 10 // Retry up to 10 times to produce the message
	producerConfig.Version = sarama.V0_10_2_0
	producerConfig.Producer.Return.Errors = true
	producerConfig.Producer.Return.Successes = true
}

func main() {
	initProducerConfig()

	go startProducers()
	select {}
}
