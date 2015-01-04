package main

import "github.com/mattn/go-pubsub"

var PubSuber = func() map[string]*pubsub.PubSub {
	res := map[string]*pubsub.PubSub{
		"Notify": pubsub.New(),
	}

	return res
}()
