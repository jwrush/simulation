package simulation

type EventCallback func(int64) []Event

// Event represents a single event in the simulation.
type Event struct {
	Time    int64
	Handler EventCallback
}

// NewEvent creates a new event.
func NewEvent(t int64, handler EventCallback) *Event {
	return &Event{
		Time:    t,
		Handler: handler,
	}
}