package main

import (
	"fmt"
	"github.com/tguangch/gtools/redis"
)

func main(){
	testNoPassward()
	//testPassward()
}

func testPassward(){
	//10.209.230.67:10816 passward:iotffan
	
	info, err := redis.Info("10.209.230.67", "10816", "iotffan")
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(info)	
}

func testNoPassward(){
	info, err := redis.Info("10.213.12.74", "12811", "")
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(info)
}