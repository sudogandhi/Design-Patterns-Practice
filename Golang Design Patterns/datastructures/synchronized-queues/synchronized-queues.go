package main

import (
	"fmt"
	"math/rand"
	"time"
)

type integerChannel chan int

const (
	messagePassStart = iota
	messageTicketStart
	messagePassEnd
	messageTicketEnd
)

type Queue struct {
	waitPass    int
	waitTicket  int
	playPass    bool
	playTicket  bool
	queuePass   integerChannel
	queueTicket integerChannel
	message     integerChannel
}

func (q *Queue) New() {
	q.message = make(integerChannel)
	q.queuePass = make(integerChannel)
	q.queueTicket = make(integerChannel)

	go func() {
		var message int
		for {
			select {
			case message = <-q.message:
				switch message {
				case messagePassStart:
					q.waitPass++
				case messagePassEnd:
					q.playPass = false
				case messageTicketStart:
					q.waitTicket++
				case messageTicketEnd:
					q.playTicket = false
				}
				if q.waitPass > 0 && q.waitTicket > 0 && !q.playPass && !q.playTicket {
					q.playPass = true
					q.playTicket = true
					q.waitTicket--
					q.waitPass--
					q.queuePass <- 1
					q.queueTicket <- 1
				}
			}
		}
	}()
}

// StartTicketIssue starts the ticket issue
func (q *Queue) StartTicketIssue() {
	q.message <- messageTicketStart
	<-q.queueTicket
}

// EndTicketIssue ends the ticket issue
func (q *Queue) EndTicketIssue() {
	q.message <- messageTicketEnd
}

func ticketIssue(q *Queue) {
	for {
		// sleep upto to 10 seconds
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
		q.StartTicketIssue()
		fmt.Println("Ticket issue starts")

		// sleep upto 2 seconds
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		fmt.Println("Ticket issue ends")
		q.EndTicketIssue()
	}
}

// StartPass ends the pass queue
func (q *Queue) StartPass() {
	q.message <- messagePassStart
	<-q.queuePass
}

// EndPass ends the pass queue
func (q *Queue) EndPass() {
	q.message <- messagePassEnd
}

func passenger(q *Queue) {
	fmt.Println("Starting the passenger Queue")
	for {
		fmt.Println("starting the processing")

		// sleep up to 10 seconds
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)

		q.StartPass()
		fmt.Println("Passenger Starts")

		// sleep up to 2 seconds
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)

		fmt.Println("Passenger Ends")
		q.EndPass()
	}
}

func main() {
	var q *Queue = &Queue{}

	fmt.Println(q)

	q.New()
	fmt.Println(q)

	var i int

	for i = 0; i < 10; i++ {
		go passenger(q)
	}

	var j int
	for j = 0; j < 5; j++ {
		go ticketIssue(q)
	}
}
