package main

type Publisher struct {
	Id string
}

func NewPublisher(Id string) *Publisher {
	return &Publisher{
		Id: Id,
	}
}

func (p *Publisher) PushToKafkaTopic(k *KafkaQueue, topic, msg string) {
	k.PublishMessage(topic, msg)
}
