package main

import (
	"fmt"
	"github.com/tguangch/gtools/stimer"
)

func main(){
	stimer.SimpleTimer {
		Period : 5,
		Task : sayHello,
	}.ExecTask()
	
}

func sayHello() {
	fmt.Println("hello go")
}