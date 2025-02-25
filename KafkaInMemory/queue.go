package main

import (
	"sync"
)

type KafkaQueue struct {
	msg map[string][]*Message

	mu sync.RWMutex
	// TODO: Register publisher, consumer and validate if correct pub, consumer are read from a topic
	// TODO: Locking if needed
}

func NewKafkaQueue() *KafkaQueue {
	return &KafkaQueue{
		msg: map[string][]*Message{},
	}
}

type Message struct {
	Content string
}

func NewKafkaMessage(msg string) *Message {
	return &Message{Content: msg}
}

func (k *KafkaQueue) RegisterATopic(topic string) {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.msg[topic] = make([]*Message, 0)
}

func (k *KafkaQueue) PublishMessage(topic, msg string) {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.msg[topic] = append(k.msg[topic], NewKafkaMessage(msg))
}

func (k *KafkaQueue) ReadFromIndex(topic string, index int, maxSize int) []*Message {
	k.mu.RLock()
	defer k.mu.RUnlock()
	messages, ok := k.msg[topic]
	if !ok {
		return nil
	}
	if index < len(messages) {
		return messages[index:min(index+maxSize, len(messages))]
	}
	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
