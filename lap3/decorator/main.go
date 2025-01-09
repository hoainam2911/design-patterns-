package main

import "fmt"

type IPizza interface {
	getPrice() int
}

type VeggieMania struct{}

func (p *VeggieMania) getPrice() int {
	return 15 
}

type BaseTopping struct {
	pizza IPizza 
}

func (b *BaseTopping) getPrice() int {
	return b.pizza.getPrice()
}

type TomatoTopping struct {
	pizza IPizza 
}

func (c *TomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice() 
	return pizzaPrice + 7 
}

type CheeseTopping struct {
	pizza IPizza 
}

func (c *CheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice() 
	return pizzaPrice + 10 
}

func main() {
	pizza := &VeggieMania{}

	pizzaWithCheese := &CheeseTopping{
		pizza: pizza,
	}

	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of VeggieMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}
