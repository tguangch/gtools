package app

import (
	"os"
	"time"
	"fmt"
)

var ch chan string = make(chan string)

func Start(){
	fmt.Println("start ...")
	
	//os.Args
	
	go echo()
	
	select {
		
	}
}

func Stop(){
	fmt.Println("stop ...")
}
func Status(){
	fmt.Println("Status")
}

func echo() {
	for {
		fmt.Println("echo ...", time.Now())
		time.Sleep(time.Second * 10)
	}
}