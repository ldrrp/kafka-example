package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/tidwall/gjson"
)

var address string
var topic string
var group string

var kafkaReader *kafka.Reader

func main() {
	config()
	// create a new context
	ctx := context.Background()
	// go produce(ctx)
	consume(ctx)
}

func config() {
	// Load env
	address = os.Getenv("ADDRESS")
	topic = os.Getenv("TOPIC")
	group = os.Getenv("GROUPID")

	if address == "" {
		log.Fatal("Empty address not allowed")
	}
	if topic == "" {
		log.Fatal("Empty topic not allowed")
	}
	if group == "" {
		log.Fatal("Empty group not allowed")
	}

	log.Printf("Loaded ENV ADDRESS: %s\n", address)
	log.Printf("Loaded ENV TOPIC: %s\n", topic)
	log.Printf("Loaded ENV GROUPID: %s\n", group)
}

func consume(ctx context.Context) {
	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
		ClientID:  "consumer-1",
		//TLS:       &tls.Config{...tls config...},
	}

	// make a new reader that consumes from topic
	kafkaReader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{address},
		GroupID:        group,
		Topic:          topic,
		MinBytes:       1,    // 1B
		MaxBytes:       10e6, // 10MB
		MaxWait:        10 * time.Second,
		StartOffset:    kafka.FirstOffset,
		Dialer:         dialer,
		CommitInterval: 1000 * time.Millisecond,
		// Logger:            Logger{"--- INFO ---"},
		// ErrorLogger:       Logger{"-- ERROR ---"},
		// HeartbeatInterval: 30 * time.Second,
		// SessionTimeout:    95 * time.Second,
	})
	defer kafkaReader.Close()

	for {
		m, err := kafkaReader.ReadMessage(ctx)
		if err != nil {
			log.Fatalf("Something went wrong reading from kafka: %s\n", err)
			break
		}

		log.Printf("message at topic/partition/offset %v/%v/%v: %s\n", m.Topic, m.Partition, m.Offset, string(m.Key))

		//Thread off processing so we dont hold up kafka
		//Could use a channel to fine tune memory and cpu usage.
		go processMessage(m.Value)

	}

}

func processMessage(message []byte) {

	filename := "/messages/visit-" + strconv.FormatInt(time.Now().UTC().UnixNano(), 10) + ".json"
	ioutil.WriteFile(filename, message, 0600)

	log.Println("User processed with ip: " + gjson.Get(string(message), "SERVER.REMOTE_ADDR").String())
}

func produce(ctx context.Context) {
	// initialize a counter
	i := 0

	// intialize the writer with the broker addresses, and the topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{address},
		Topic:   topic,
	})

	for {
		// each kafka message has a key and value. The key is used
		// to decide which partition (and consequently, which broker)
		// the message gets published on
		err := w.WriteMessages(ctx, kafka.Message{
			Key: []byte(strconv.Itoa(i)),
			// create an arbitrary message payload for the value
			Value: []byte("this is message" + strconv.Itoa(i)),
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}

		// log a confirmation once the message is written
		fmt.Println("writes:", i)
		i++
		// sleep for a second
		time.Sleep(time.Second)
	}
}
