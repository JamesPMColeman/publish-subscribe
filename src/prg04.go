/************************************************************************
*																		*
*																		*
*			James Coleman												*
*			CS 3210														*
*			Project 5													*
*			November 30th												*
*																		*
*																		*
************************************************************************/


	   /*
		*	Thanks to Eli Bendersky @ eli.thegreenplace.net	
		*
		*	Facts from:
		*		New York Times - Not all rivers reach the sea
		*		Business Insider - China's Monster Three Gorges Dam ...
		*		Kickassfacts.com - 25 kickass interesting facts ...
		*		Amusingplanet.com - A Stream that swallows people
		*		Kiwi.com - 15 unusual facts about world cities	
		*
		*/

package main

import (
	"sync";
	"fmt";
	"time";
	"math/rand";
)

type PubSub struct {
	mu sync.Mutex
	topics map[string][]chan string
}

var wg sync.WaitGroup

// Done: writes the given message on all the channels associated with the given topic
func (ps PubSub) publish(topic string, fact string) {
	//fmt.Println("Publish")

	ps.mu.Lock()

	for _, channel := range ps.topics[topic] {
		go func(c chan string) {
			c <- fact
		}(channel)
	}
	ps.mu.Unlock()
}

// Done: sends messages taken from a given array of message, one at a time and at random intervals, to all topic subscribers
func publisher(ps PubSub, topic string, facts []string) {
	//fmt.Println("Publisher")

	time.Sleep(5 * time.Second)

	for _, fact := range facts{
		//fmt.Println("****** " + fact)
		time.Sleep(time.Duration(rand.Intn(7)) * time.Second)
		ps.publish(topic, fact)
	}

	time.Sleep(5 * time.Second)
	wg.Done()
}

// Done: creates and returns a new channel on a given topic, updating the PubSub struct
func (ps PubSub) subscribe(topic string) <-chan string {
	//fmt.Println("Subscribe: " + topic)

	ps.mu.Lock()

	factChannel := make(chan string)
	//fmt.Println(factChannel)
	ps.topics[topic] = append(ps.topics[topic], factChannel)

	ps.mu.Unlock()
	return factChannel
}

// Done: reads and displays all messages received from a particular topic
func subscriber(ps PubSub, name string, topic string) {
	//fmt.Println("Subscriber")

	factsChannel := ps.subscribe(topic)
	// fmt.Printf("\t%T\n", factsChannel)
	// fmt.Println(len(factsChannel))

	for {
		//fmt.Println(factsChannel)
		//fmt.Println(len(factsChannel))
		//fmt.Printf("%T\n", factsChannel)

		if fact, ready := <-factsChannel; ready {
			fmt.Println(name + " received: " + fact)
		} else {
			break
		}
	}
}

func main() {

	// Done: create the ps struct
	ps := PubSub{topics: make(map[string][]chan string)}

	// Done: create the arrays of messages to be sent on each topic
	riverFacts := []string{
		"The Colorado river has not reached the sea since 1998",
		"Three Gorges Dam on the Yangtze river in china, holds so much water that it slowed the rotation of the Earth ever so slightly",
		"The Bolton Stride in England is only about six feet across but it is so deep, know one can say where the bottom is. Because of this fact the stride is very dangerous",
	}

	cityFacts := []string{
		"Paris has no stop signs",
		"Seattle has more households with pets than with children",
		"Mexico City is built on lake Texcoco, as a result it is constantly sinking",
		"In Calcutta more people use bikes then cars to get around",
	}

	desertFacts := []string{
		"Saying 'Sahara desert' is redundant. The word sahara means desert in Arabic",
		"The largest desert in the world is Antarctica",
		"Sand dunes 'sing' as the shift",
	}

	// Done: set wait group to 2 (# of publishers)
	wg.Add(2)

	fmt.Println("\n\n>>>>>>>>>>> Welcome <<<<<<<<<<<<\n\n")

	// Done: create the publisher goroutines
	go publisher(ps, "River Facts", riverFacts)
	go publisher(ps, "City Facts", cityFacts)
	go publisher(ps, "Desert Facts", desertFacts)

	// Done: create the subscriber goroutines
	go subscriber(ps, "Anna", "River Facts")
	go subscriber(ps, "Anna", "Desert Facts")
	go subscriber(ps, "Ben", "River Facts")
	go subscriber(ps, "Ben", "City Facts")
	go subscriber(ps, "Ben", "Desert Facts")
	go subscriber(ps, "Cal", "Desert Facts")
	go subscriber(ps, "Danny", "River Facts")
	go subscriber(ps, "Danny", "City Facts")

	// Done: wait for all publishers to be done
	wg.Wait()
}
