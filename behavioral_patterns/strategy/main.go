package main

import "fmt"

type RouteStrategy interface {
	BuildRoute(start, end string) string
}

type ConcreteStrategyRoad struct{}

func (r *ConcreteStrategyRoad) BuildRoute(start, end string) string {
	return fmt.Sprintf("Tuyến đường bằng ô tô từ %s đến %s.", start, end)
}

type ConcreteStrategyWalking struct{}

func (w *ConcreteStrategyWalking) BuildRoute(start, end string) string {
	return fmt.Sprintf("Tuyến đường đi bộ từ %s đến %s.", start, end)
}

type ConcreteStrategyPublicTransport struct{}

func (pt *ConcreteStrategyPublicTransport) BuildRoute(start, end string) string {
	return fmt.Sprintf("Tuyến đường bằng phương tiện công cộng từ %s đến %s.", start, end)
}

type Navigator struct {
	strategy RouteStrategy
}

func (n *Navigator) SetStrategy(strategy RouteStrategy) {
	n.strategy = strategy
}

func (n *Navigator) BuildRoute(start, end string) {
	if n.strategy == nil {
		fmt.Println("Chưa có chiến lược định tuyến nào được chọn.")
		return
	}
	route := n.strategy.BuildRoute(start, end)
	fmt.Println(route)
}

func main() {
	navigator := &Navigator{}
	navigator.SetStrategy(&ConcreteStrategyRoad{})
	navigator.BuildRoute("Hà Nội", "Hải Phòng")
	navigator.SetStrategy(&ConcreteStrategyWalking{})
	navigator.BuildRoute("Hà Nội", "Hải Phòng")
	navigator.SetStrategy(&ConcreteStrategyPublicTransport{})
	navigator.BuildRoute("Hà Nội", "Hải Phòng")
}
