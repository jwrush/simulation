
package simulation

import (
	"container/heap"
	"errors"
	"fmt"
)

// EventQueue manages a list of events sorted by time.
type EventQueue []*Event

func (eq EventQueue) Len() int           { return len(eq) }
func (eq EventQueue) Less(i, j int) bool { return eq[i].Time < eq[j].Time }
func (eq EventQueue) Swap(i, j int)      { eq[i], eq[j] = eq[j], eq[i] }

func (eq *EventQueue) Push(x interface{}) {
	*eq = append(*eq, x.(*Event))
}

func (eq *EventQueue) Pop() interface{} {
	old := *eq
	n := len(old)
	event := old[n-1]
	*eq = old[0 : n-1]
	return event
}

// Simulation manages the event queue and simulation time.
type Simulation struct {
	Queue   EventQueue
	SimTime int64
}

// NewSimulation creates a new simulation.
func NewSimulation(startTime int64) *Simulation {
	return &Simulation{
		Queue:   make(EventQueue, 0),
		SimTime: startTime,
	}
}

// Schedule adds a new event to the queue.
func (s *Simulation) Schedule(event *Event) error {
	if event.Time <= s.SimTime {
		return errors.New("event time must be in the future")
	}
	heap.Push(&s.Queue, event)
	return nil
}

// Run processes events in the queue in order.
func (s *Simulation) Run() {
	for s.Queue.Len() > 0 {
		event := heap.Pop(&s.Queue).(*Event)
		s.SimTime = event.Time
		newEvents := event.Handler(s.SimTime)

		for _, e := range newEvents {
			err := s.Schedule(&e)
			if err != nil {
				fmt.Println("Error scheduling event:", err)
			}
		}
	}
}
