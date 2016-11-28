package main
import (
	"github.com/tguangch/gtools/stats/machine"
	"log"
	"encoding/json"
)

func main(){
	topItems := machine.Top()
	log.Println(topItems)
	
	b, err := json.Marshal(topItems)
	if err != nil {
		log.Fatalln("error,", err)
	}
	log.Println(string(b))
}
