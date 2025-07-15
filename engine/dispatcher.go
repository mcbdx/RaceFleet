package engine

import (
	"github.com/mcbdx/RaceFleet/car"
)

// Dispatcher - take a list of cars 
// kick off the race by launching each car in a goroutine
// collect the results of each car
// return the result in a channel 
// determine winner based on the results
// Record Stats and return the stats

// start race 
