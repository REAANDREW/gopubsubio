package gopubsubio

import (
	"fmt"
	"sync"
	"testing"
)

func Test_PublishingAMessageToMultipleSubscribers(t *testing.T) {

	var wg sync.WaitGroup

	publisher := NewPublisher()

	wg.Add(1)
	subscriber1 := NewSubscriber(func(value interface{}) {
		fmt.Printf("Subscriber 1 received %v\n", value)
		wg.Done()
	})

	wg.Add(1)
	subscriber2 := NewSubscriber(func(value interface{}) {
		fmt.Printf("Subscriber 2 received %v\n", value)
		wg.Done()
	})

	publisher.Subscribe("foobar", subscriber1)
	publisher.Subscribe("foobar", subscriber2)
	publisher.Publish("foobar", 1)

	wg.Wait()
}
