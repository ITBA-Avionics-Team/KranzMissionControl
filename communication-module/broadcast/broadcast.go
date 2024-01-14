package broadcast

import (
	"errors"
	"fmt"
)

// TODO: Implement this with a mutex to prevent race conditions
type Broadcast[T any] struct {
	subscribers map[int]chan T
	nextId int
}

func NewBroadcast[T any]() *Broadcast[T] {
	return &Broadcast[T]{
		subscribers: make(map[int]chan T),
	}
}

func(broadcast * Broadcast[T]) Subscribe() (id int, subscription chan T){
	newSubscription := make(chan T)
	newSubscriptionId := broadcast.nextId
	broadcast.subscribers[broadcast.nextId] = newSubscription
	broadcast.nextId += 1
	return newSubscriptionId, newSubscription
}

func(broadcast * Broadcast[T]) Unubscribe(subscriptionId int) (err error){
	ch, ok := broadcast.subscribers[subscriptionId]
	if ok {
		close(ch)
		delete(broadcast.subscribers, subscriptionId)
		fmt.Println("Unsubscribing....")
		return nil
	} else {
		return errors.New("No subscription found for given id " + string(subscriptionId))
	}
}

func(broadcast * Broadcast[T]) SendBroadcast(message T) {
	for _, ch := range broadcast.subscribers {
		ch <- message
	}
}