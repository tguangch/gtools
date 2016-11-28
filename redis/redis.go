package redis

import (
	"fmt"
	//"os"
	"flag"
)

var (
	host = flag.String("host", "hhh", "host name")
)

func Go(){
	//flag.Parse()
	fmt.Println(flag.Parsed())
//	fmt.Println(flag.CommandLine)
//	flag.Parse()
//	
//	
//	
//	fmt.Println(os.Args)
//	fmt.Println(flag.CommandLine)
	fmt.Println(*host)
}

