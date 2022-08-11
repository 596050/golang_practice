package main

import (
	. "section_3/producer_consumer"
)

func main() {
	// var wg sync.WaitGroup
	producerConsumer()

}

// Race conditions -
// - When two different routines access the same data
// - have no information on when they start and end, both change the data

// Need to test for race conditions
