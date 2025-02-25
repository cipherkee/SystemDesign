package main

import (
	"fmt"
	"sync"
	"time"
)

/*
The queue should be in-memory and should not require access to the file system.
There can be multiple topics in the queue.
A (string) message can be published on a topic by a producer/publisher and consumers/subscribers can subscribe to the topic to receive the messages.
There can be multiple producers and consumers.
A producer can publish to multiple topics.
A consumer can listen to multiple topics.
The consumer should print "<consumer_id> received <message>" on receiving the message.
The queue system should be multi-threaded, i.e., messages can be produced or consumed in parallel by different producers/consumers.
*/

func main() {

	var (
		topic1 = "topic1"
		wg     sync.WaitGroup
	)

	k := NewKafkaQueue()
	// 1 topic
	k.RegisterATopic(topic1)

	// 2 publishers
	p1 := NewPublisher("p1")
	p2 := NewPublisher("p2")

	// 2 consumers
	c1 := NewConsumer("c1")
	c2 := NewConsumer("c2")

	c1.RegisterToTopic(topic1, 3)
	c2.RegisterToTopic(topic1, 3)
	go func() {
		wg.Add(1)
		c1.Runnable(k)
	}()
	go func() {
		wg.Add(1)
		c2.Runnable(k)
	}()

	time.Sleep(2 * time.Second)
	go func() {
		wg.Add(1)
		publishMessageEveryTSec(k, p1, topic1, 5)
	}()
	go func() {
		wg.Add(1)
		publishMessageEveryTSec(k, p2, topic1, 5)
	}()

	wg.Wait()
}

func publishMessageEveryTSec(k *KafkaQueue, p *Publisher, topic string, t int) {
	count := 0
	for {
		time.Sleep(time.Duration(t) * time.Second)
		p.PushToKafkaTopic(k, topic, fmt.Sprintf("message:%v:%vs::%v", p.Id, t, count))
		count += t
	}
}
