package simulation

import (
	"testing"
)

func TestSimulation(t *testing.T) {
	startTime := int64(0)
	sim := NewSimulation(startTime)

	var event1Triggered, event2Triggered bool

	sim.Schedule(NewEvent(
		startTime + 1,
		func(_ int64) []Event { event1Triggered = true; return nil },
	))
	sim.Schedule(NewEvent(
		startTime + 2,
		func(_ int64) []Event { event2Triggered = true; return nil },
	))

	sim.Run()

	if !event1Triggered {
		t.Errorf("Event 1 was not triggered")
	}
	if !event2Triggered {
		t.Errorf("Event 2 was not triggered")
	}
}

func TestSimulationChildEvents(t *testing.T) {
	startTime := int64(0)
	sim := NewSimulation(startTime)

	var event1Triggered, event2Triggered bool
	var event1Time, event2Time int64

	sim.Schedule(NewEvent(
		startTime + 1,
		func(time int64) []Event {
			event1Triggered = true
			event1Time = time
			return []Event{
				{
					Time: time + 2,
					Handler: func(time int64) []Event {
						event2Triggered = true
						event2Time = time
						return nil
					},
				},
			}
		},
	))

	sim.Run()

	if !event1Triggered {
		t.Errorf("Event 1 was not triggered")
	}
	if !event2Triggered {
		t.Errorf("Event 2 was not triggered")
	}
	if event1Time != 1 {
		t.Errorf("Event 1 was triggered at the wrong time: %d", event1Time)
	}
	if event2Time != 3 {
		t.Errorf("Event 2 was triggered at the wrong time: %d", event2Time)
	}
}
