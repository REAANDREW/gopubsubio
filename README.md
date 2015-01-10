# gopubsubio
Simple Pub/Sub Package for golang

Create a new Publisher

```go
import (
    "github.com/REAANDREW/gopubsubio"
)

publisher := gopubsubio.NewPublisher()
```

Create a couple of Subscribers

```go
subscriber1 := gopubsubio.NewSubscriber(func(value interface{}) {
    fmt.Printf("Subscriber 1 received %v\n", value)
})

subscriber2 := gopubsubio.NewSubscriber(func(value interface{}) {
    fmt.Printf("Subscriber 2 received %v\n", value)
})
```

Add the subscribers to the publisher for a specific `key`

```go
publisher.Subscribe("foobar", subscriber1)
publisher.Subscribe("foobar", subscriber2)
```

Publish an event to notify the subscribers

```go
publisher.Publish("foobar", 1)
```
