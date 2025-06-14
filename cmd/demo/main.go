package main

import (
	"fmt"
	"log"

	"simulation"
)

func main() {
	fmt.Println("Starting Discrete Event Simulation Demo")
	
	// Create a new simulation starting at time 0
	sim := simulation.NewSimulation(0)
	
	// Example event handler that prints a message and schedules a follow-up event
	exampleHandler := func(currentTime int64) []simulation.Event {
		fmt.Printf("Event executed at time %d\n", currentTime)
		
		// Return empty slice for no follow-up events, or schedule new events
		return []simulation.Event{}
	}
	
	// Recursive event handler that demonstrates chaining events
	var recursiveHandler simulation.EventCallback
	recursiveHandler = func(currentTime int64) []simulation.Event {
		fmt.Printf("Recursive event at time %d\n", currentTime)
		
		// Stop recursion after time 100
		if currentTime < 100 {
			nextEvent := simulation.Event{
				Time:    currentTime + 10,
				Handler: recursiveHandler,
			}
			return []simulation.Event{nextEvent}
		}
		return []simulation.Event{}
	}
	
	// Create some initial events
	event1 := simulation.NewEvent(5, exampleHandler)
	event2 := simulation.NewEvent(10, recursiveHandler)
	event3 := simulation.NewEvent(25, func(t int64) []simulation.Event {
		fmt.Printf("One-time event at time %d\n", t)
		return []simulation.Event{}
	})
	
	// Schedule the events
	if err := sim.Schedule(event1); err != nil {
		log.Fatalf("Failed to schedule event1: %v", err)
	}
	
	if err := sim.Schedule(event2); err != nil {
		log.Fatalf("Failed to schedule event2: %v", err)
	}
	
	if err := sim.Schedule(event3); err != nil {
		log.Fatalf("Failed to schedule event3: %v", err)
	}
	
	fmt.Println("Running simulation...")
	
	// Run the simulation
	sim.Run()
	
	fmt.Printf("Simulation completed. Final time: %d\n", sim.SimTime)
}