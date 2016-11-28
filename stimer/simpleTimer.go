package stimer

import (
	"time"
	"log"
)

type SimpleTimer struct {
	Period	time.Duration
	Task	func()
}

func (timer SimpleTimer) ExecTask(){
	for {
		time.Sleep(timer.Period * time.Second)
		log.Println("task go")
		go timer.Task()
	}
}
