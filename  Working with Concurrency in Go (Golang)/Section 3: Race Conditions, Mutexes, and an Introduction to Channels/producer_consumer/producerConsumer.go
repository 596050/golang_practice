package producer_consumer

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

// anything of type Producer has access to this
func (p *Producer) Close() error {
	ch := make(chan error)
	// assign a channel
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("received order #%d!\n", pizzaNumber)
		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++
		fmt.Printf("Making pizza #%d. It will take %d seconds.... \n", pizzaNumber, delay)
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d!", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quit while making pizza #%d!", pizzaNumber)

		} else {
			success = true
			msg = fmt.Sprintf("Pizza order #%d is ready", pizzaNumber)
		}
		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}
		return &p
	}

	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzeria(pizzaMaker *Producer) {
	// keep track of each pizza we are making
	var i int = 0

	// run forever or until we receive a quit notification

	// try to make pizza

	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
		}
	}
}

func producerConsumer() {
	// seed random number generator
	rand.Seed(time.Now().UnixNano())

	// print out message
	color.Cyan("The Pizzaria is open for business")
	color.Cyan("---------------------------------")

	// create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run the producer in the background
	go pizzeria(pizzaJob)

	// create and run consumer

	// print our the ending message

	return
}

// channels are far more useful than mutexes and waitgroups
// they allow you to exchange data to another go routine
