package main

import (
	"invasion/world"
	"flag"
	"os"
	"fmt"

)


var num_aliens int
var file string

func init() {
	flag.IntVar(&num_aliens, "aliens", 8, "Number of aliens")
	flag.StringVar(&file, "file", "testmap", "Map containing cities and their neighbors") 
}


func main() {



	
	flag.Parse()

	
	WorldMap, e := world.DecodeMap(file)

	if e != nil {
		fmt.Printf("Couldn't open file %s.", e)
		os.Exit(-1)
	}
	
	WorldMap.InitAliens(num_aliens)
	Srv := world.MakeServer(WorldMap) 



	
	go Srv.Handler()
	go Srv.SpawnWorkers()
	<- Srv.Sig 

}
