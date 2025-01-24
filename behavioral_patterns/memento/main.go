package main

import "fmt"

type Memento struct {
	state string
}

type Originator struct {
	state string
}

func (o *Originator) SetState(state string) {
	o.state = state
	fmt.Println("State set to:", state)
}

func (o *Originator) Save() *Memento {
	fmt.Println("Saving state:", o.state)
	return &Memento{state: o.state}
}

func (o *Originator) Restore(m *Memento) {
	o.state = m.state
	fmt.Println("State restored to:", o.state)
}

type Caretaker struct {
	history []*Memento
}

func (c *Caretaker) AddMemento(m *Memento) {
	c.history = append(c.history, m)
	fmt.Println("Memento added to history.")
}

func (c *Caretaker) GetMemento(index int) *Memento {
	if index >= 0 && index < len(c.history) {
		fmt.Println("Fetching memento at index:", index)
		return c.history[index]
	}
	fmt.Println("Invalid index.")
	return nil
}

func main() {
	originator := &Originator{}
	caretaker := &Caretaker{}

	originator.SetState("State1")
	caretaker.AddMemento(originator.Save())

	originator.SetState("State2")
	caretaker.AddMemento(originator.Save())

	originator.SetState("State3")

	originator.Restore(caretaker.GetMemento(1))
	originator.Restore(caretaker.GetMemento(0))
}
