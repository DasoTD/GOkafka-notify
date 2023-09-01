package main

import (
    "time"

    // "github.com/Shopify/sarama"

    "github.com/IBM/sarama"
)

func main1() {
    config := sarama.NewConfig()
    config.Producer.RequiredAcks = sarama.WaitForLocal
    config.Producer.Compression = sarama.CompressionSnappy
    config.Producer.Flush.Frequency = 500 * time.Millisecond

    producer, err := sarama.NewAsyncProducer([]string{"kafka-broker:9092"}, config)
    if err != nil {
        // Handle error
        return
    }
    defer producer.Close()

    // go func() {
    //     for err := range producer.Errors() {
    //         // Handle error
	// 		return err
    //     }
    // }()

    go func() {
        for _ = range producer.Successes() {
            // Message successfully sent
        }
    }()

    message := &sarama.ProducerMessage{
        Topic: "notifications",
        Key:   sarama.StringEncoder("key"),
        Value: sarama.StringEncoder("your-message"),
    }

    producer.Input() <- message

    // Optionally wait for the message to be sent
    // time.Sleep(1 * time.Second)
}
