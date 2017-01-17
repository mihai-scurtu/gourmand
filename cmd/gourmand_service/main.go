package main

import (
	"time"

	"github.com/mihai-scurtu/gourmand"
)

func main() {
	app := gourmand.NewApp()
	sched := newScheduler(app, 10*time.Minute, 1*time.Hour)

	sched.run()
}
