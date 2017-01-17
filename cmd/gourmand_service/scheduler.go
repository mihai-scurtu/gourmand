package main

import (
	"log"
	"time"

	"github.com/mihai-scurtu/gourmand"
)

type scheduler struct {
	app         *gourmand.App
	ticker      <-chan time.Time
	waiting     bool
	waitTimeout time.Duration
	q           chan interface{}
}

func newScheduler(app *gourmand.App, timeout time.Duration, waitTimeout time.Duration) *scheduler {
	return &scheduler{
		app:         app,
		ticker:      time.Tick(timeout),
		waiting:     false,
		waitTimeout: waitTimeout,
		q:           make(chan interface{}),
	}
}

func (s scheduler) run() {
loop:
	for {
		s.waiting = false

		log.Println("Running process...")

		err := s.app.Run()
		if err != nil {
			switch err {
			case gourmand.NoItemsError:
				log.Printf("No items found, waiting %s.", s.waitTimeout.String())
				s.waiting = true
			default:
				log.Printf("Error: %s", err.Error())
			}
		}

		log.Println("Done.")

		select {
		case <-s.timing():

		case <-s.q:
			break loop
		}
	}
}

func (s scheduler) timing() <-chan time.Time {
	if s.waiting {
		return time.After(s.waitTimeout)
	}

	return s.ticker
}

func (s scheduler) quit() {
	s.q <- nil
}
