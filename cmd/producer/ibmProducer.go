package main

import (
    "time"

    "github.com/IBM/sarama"
)

func mainx() {
    config := sarama.NewConfig()
    config.Producer.RequiredAcks = sarama.WaitForLocal
    config.Producer.Compression = sarama.CompressionSnappy
    config.Producer.Flush.Frequency = 500 * time.Millisecond

    // Set additional configuration specific to IBM Cloud Event Streams, if necessary
    config.Net.SASL.Enable = true
    config.Net.SASL.User = "your-username"
    config.Net.SASL.Password = "your-password"

    producer, err := sarama.NewAsyncProducer([]string{"kafka-broker:9093"}, config)
    if err != nil {
        // Handle error
        return
    }
    defer producer.Close()

    // go func() {
    //     for err := range producer.Errors() {
    //         // Handle error
    //     }
    // }()

    go func() {
        for _ = range producer.Successes() {
            // Message successfully sent
        }
    }()

    message := &sarama.ProducerMessage{
        Topic: "your-topic",
        Key:   sarama.StringEncoder("key"),
        Value: sarama.StringEncoder("your-message"),
    }

    producer.Input() <- message

    // Optionally wait for the message to be sent
    // time.Sleep(1 * time.Second)
}
