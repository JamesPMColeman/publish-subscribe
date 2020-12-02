/************************************************************************
*																		*
*																		*
*			James Coleman												*
*			CS 3210														*
*			Project 5													*
*			November 30th												*
*																		*
************************************************************************/


	   /*
		*		Thanks to youthincmag.com for instrument facts
		*
		*
		*
		*/

package main

import (
	"sync";
	"fmt"
)

type PubSub struct {
	mu sync.Mutex
	topics map[string][]chan string
}

var wg sync.WaitGroup

// TODO: creates and returns a new channel on a given topic, updating the PubSub struct
func (ps PubSub) subscribe(topic string) chan string {
	ps.mu.Lock()
	fmt.Println("subscribe")
	ps.topics[topic] = append(ps.topics[topic], "Yo")
	ps.mu.Unlock()
	return nil
}

// TODO: writes the given message on all the channels associated with the given topic
func (ps PubSub) publish(topic string, facts string) {
	fmt.Println("Publish")
}

// TODO: sends messages taken from a given array of message, one at a time and at random intervals, to all topic subscribers
func publisher(ps PubSub, topic string, facts[]string) {
	fmt.Println("Publisher")
}

// TODO: reads and displays all messages received from a particular topic
func subscriber(ps PubSub, name string, topic string) {
	factsChannel := ps.Subscribe(topic)
	fmt.Println("Subscriber")
}

func main() {

	// TODO: create the ps struct
	ps := PubSub{topics: make(map[string][]chan string)}

	// Done: create the arrays of messages to be sent on each topic
	riverFacts := []string{
		"The Colorado river has not reached the sea since 1998",
		"Three Gorges Dam on the Yangtze river in china, holds so much waterthat it slowed the rotation of the Earth ever so slightly",
		"The Bolton Stride in England is only about six feet accross but it is so deep, know one can say where the bottom is. Becuase of this fact the stride is very dangerous",
	}

	cityFacts := []string{
		"Paris has no stop signs",
		"Seattle has more hoseholds with pets than with children",
		"Mexico City is built on lake Texcoco, as a result it is constantly sinking",
		"In Calcutta more people use bikes then cars to get around",
	}

	desertFacts := []string{
		"Saying 'Sahara desert' is redundent. The word sahara means desert in Arabic",
		"The largest desert in the world is Antarctica",
		"Sand dunes 'sing' as the shift",
	}

	// Done: set wait group to 2 (# of publishers)
	wg.Add(2)

	// Done: create the publisher goroutines
	go publisher(ps, "River Facts", riverFacts)
	go publisher(ps, "City Facts", cityFacts)
	go publisher(ps, "Desert Facts", desertFacts)

	// Done: create the subscriber goroutines
	go subscriber(ps, "Anna", riverFacts)
	go subscriber(ps, "Anna", desertFacts)
	go subscriber(ps, "Ben", riverFacts)
	go subscriber(ps, "Ben", cityFacts)
	go subscriber(ps, "Ben", desertFacts)
	go subscriber(ps, "Cal", desertFacts)
	go subscriber(ps, "Danny", riverFacts)
	go subscriber(ps, "Danny", cityFacts)

	// Done: wait for all publishers to be done
	wg.Wait()
}
