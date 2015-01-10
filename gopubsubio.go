package gopubsubio

import (
	"time"
)

type Subscriber interface {
	Notify(messgae interface{})
}

type Publisher interface {
	Subscribe(key string, subscriber Subscriber)
	Publish(key string, message interface{})
}

type PublishDelegate func(message interface{})

func NewSubscriber(publishDelegate PublishDelegate) (subscriber Subscriber) {
	defaultSubscriber := &DefaultSubscriber{make(chan interface{})}
	go defaultSubscriber.handlePublications(publishDelegate)
	return defaultSubscriber
}

type DefaultSubscriber struct {
	channel chan interface{}
}

func (subscriber *DefaultSubscriber) handlePublications(publishDelegate PublishDelegate) {
	for {
		select {
		case message, ok := <-subscriber.channel:
			if ok {
				publishDelegate(message)
			} else {
				subscriber.channel = nil
				return
			}
		default:
			time.Sleep(1)
		}
	}
}

func (subscriber *DefaultSubscriber) Notify(message interface{}) {
	subscriber.channel <- message
}

type DefaultPublisher struct {
	subscribers map[string][]Subscriber
}

func (publisher *DefaultPublisher) Subscribe(key string, subscriber Subscriber) {
	if subscriberList := publisher.subscribers[key]; subscriberList == nil {
		publisher.subscribers[key] = []Subscriber{}
	}

	publisher.subscribers[key] = append(publisher.subscribers[key], subscriber)
}

func (publisher *DefaultPublisher) Publish(key string, message interface{}) {
	subscriberList := publisher.subscribers[key]
	for _, subscriber := range subscriberList {
		if subscriber != nil {
			subscriber.Notify(message)
		}
	}
}

func NewPublisher() (publisher Publisher) {
	publisher = &DefaultPublisher{make(map[string][]Subscriber)}
	return
}
