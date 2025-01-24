package main

import "fmt"

// Observer defines the interface for objects that should be notified of changes.
type Observer interface {
	Update(event string)
}

// Subject defines the interface for objects that manage subscribers.
type Subject interface {
	Subscribe(observer Observer)
	Unsubscribe(observer Observer)
	Notify(event string)
}

// Store represents the publisher (subject) in the Observer pattern.
type Store struct {
	observers []Observer
}

func (s *Store) Subscribe(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Store) Unsubscribe(observer Observer) {
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *Store) Notify(event string) {
	for _, observer := range s.observers {
		observer.Update(event)
	}
}

// Customer represents a subscriber in the Observer pattern.
type Customer struct {
	Name string
}

func (c *Customer) Update(event string) {
	fmt.Printf("%s received notification: %s\n", c.Name, event)
}

func main() {
	// Create a store (subject)
	store := &Store{}

	// Create customers (observers)
	customer1 := &Customer{Name: "Alice"}
	customer2 := &Customer{Name: "Bob"}

	// Subscribe customers to the store
	store.Subscribe(customer1)
	store.Subscribe(customer2)

	// Notify customers of a new product arrival
	store.Notify("New iPhone is now available!")

	// Unsubscribe one customer
	store.Unsubscribe(customer1)

	// Notify customers again
	store.Notify("Special discount on the new iPhone!")
}
