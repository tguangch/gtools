package main

import (
	"time"
	"fmt"
)

func main(){
	
	go echo()
	
	select {
		
	}
}

func echo(){
	for {
		time.Sleep(time.Duration(5) * time.Second)
		go fmt.Println("hello go ...")
	}
	
}
