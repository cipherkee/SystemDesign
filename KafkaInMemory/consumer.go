package main

import (
	"fmt"
	"time"
)

type Consumer struct {
	Id     string
	Topics map[string]*ConsumerMetadata
}

type ConsumerMetadata struct {
	lastReadIndex int
	maxSize       int // Can be configured here, or at common consumer layer
}

func NewConsumer(id string) *Consumer {
	return &Consumer{
		Id:     id,
		Topics: map[string]*ConsumerMetadata{},
	}
}

func (c *Consumer) RegisterToTopic(topicName string, maxSizeOfRead int) {
	c.Topics[topicName] = &ConsumerMetadata{lastReadIndex: 0, maxSize: maxSizeOfRead}
}

func (c *Consumer) Runnable(k *KafkaQueue) {
	// Every 2 secs, check for 3 messages in the kafka queue
	for {
		time.Sleep(2 * time.Second)
		for name, data := range c.Topics {
			messages := k.ReadFromIndex(name, data.lastReadIndex+1, data.maxSize)
			c.ProcessMessage(messages)
			data.lastReadIndex += len(messages)
		}
	}
}

func (c *Consumer) ProcessMessage(messages []*Message) {
	for i := 0; i < len(messages); i++ {
		fmt.Println(fmt.Sprintf("%v:%v", c.Id, messages[i].Content))
	}
}
