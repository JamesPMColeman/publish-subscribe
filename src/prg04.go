/*
 * CS3210 - Principles of Programming Languages - Fall 2020
 * Instructor: Thyago Mota
 * Description: Prg04 - Publish Subscribe Simulation
 * Student(s) Name(s):
 */

package main

import (
	"sync"
)

type PubSub struct {
	mu sync.Mutex
	topics map[string][]chan string
}

var wg sync.WaitGroup

// TODO: creates and returns a new channel on a given topic, updating the PubSub struct
func (ps PubSub) subscribe(topic string) chan string {
	ps.mu.Lock()
	ps.topics[topic]
	
	ps.mu.Unlock()
	return nil
}

// TODO: writes the given message on all the channels associated with the given topic
func (ps PubSub) publish(topic string, msg string) {

}

// TODO: sends messages taken from a given array of message, one at a time and at random intervals, to all topic subscribers
func publisher(ps PubSub, topic string, msgs[]string) {

}

// TODO: reads and displays all messages received from a particular topic
func subscriber(ps PubSub, name string, topic string) {

}

func main() {

	// TODO: create the ps struct
	ps := PubSub{topics: make(map[string][]chan string)}

	// TODO: create the arrays of messages to be sent on each topic


	// TODO: set wait group to 2 (# of publishers)

	// TODO: create the publisher goroutines


	// TODO: create the subscriber goroutines


	// TODO: wait for all publishers to be done
}
