package main

import "fmt"

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type IterableCollection interface {
	CreateIterator() Iterator
}

type ConcreteCollection struct {
	elements []string
}

func (c *ConcreteCollection) CreateIterator() Iterator {
	return &ConcreteIterator{
		collection: c,
		index:      0,
	}
}

type ConcreteIterator struct {
	collection *ConcreteCollection
	index      int
}

func (it *ConcreteIterator) HasNext() bool {
	return it.index < len(it.collection.elements)
}

func (it *ConcreteIterator) Next() interface{} {
	if it.HasNext() {
		element := it.collection.elements[it.index]
		it.index++
		return element
	}
	return nil
}

func main() {
	collection := &ConcreteCollection{
		elements: []string{"Element1", "Element2", "Element3"},
	}

	iterator := collection.CreateIterator()

	for iterator.HasNext() {
		element := iterator.Next()
		fmt.Println(element)
	}
}
