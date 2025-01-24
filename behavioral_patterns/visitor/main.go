package main

import "fmt"

type Shape interface {
	Accept(visitor Visitor)
}

type Dot struct {
	ID int
	X, Y int
}

func (d *Dot) Accept(visitor Visitor) {
	visitor.VisitDot(d)
}

type Circle struct {
	ID     int
	X, Y   int
	Radius int
}

func (c *Circle) Accept(visitor Visitor) {
	visitor.VisitCircle(c)
}

type Rectangle struct {
	ID     int
	X, Y   int
	Width  int
	Height int
}

func (r *Rectangle) Accept(visitor Visitor) {
	visitor.VisitRectangle(r)
}

type CompoundShape struct {
	ID     int
	Shapes []Shape
}

func (cs *CompoundShape) Accept(visitor Visitor) {
	visitor.VisitCompoundShape(cs)
}

type Visitor interface {
	VisitDot(d *Dot)
	VisitCircle(c *Circle)
	VisitRectangle(r *Rectangle)
	VisitCompoundShape(cs *CompoundShape)
}

type XMLExportVisitor struct{}

func (x *XMLExportVisitor) VisitDot(d *Dot) {
	fmt.Printf("<dot id=\"%d\" x=\"%d\" y=\"%d\" />\n", d.ID, d.X, d.Y)
}

func (x *XMLExportVisitor) VisitCircle(c *Circle) {
	fmt.Printf("<circle id=\"%d\" x=\"%d\" y=\"%d\" radius=\"%d\" />\n", c.ID, c.X, c.Y, c.Radius)
}

func (x *XMLExportVisitor) VisitRectangle(r *Rectangle) {
	fmt.Printf("<rectangle id=\"%d\" x=\"%d\" y=\"%d\" width=\"%d\" height=\"%d\" />\n", r.ID, r.X, r.Y, r.Width, r.Height)
}

func (x *XMLExportVisitor) VisitCompoundShape(cs *CompoundShape) {
	fmt.Printf("<compound_shape id=\"%d\">\n", cs.ID)
	for _, shape := range cs.Shapes {
		shape.Accept(x)
	}
	fmt.Println("</compound_shape>")
}

func main() {
	dot := &Dot{ID: 1, X: 10, Y: 20}
	circle := &Circle{ID: 2, X: 15, Y: 25, Radius: 10}
	rectangle := &Rectangle{ID: 3, X: 5, Y: 5, Width: 20, Height: 15}
	compound := &CompoundShape{
		ID: 4,
		Shapes: []Shape{dot, circle, rectangle},
	}

	exportVisitor := &XMLExportVisitor{}

	dot.Accept(exportVisitor)
	circle.Accept(exportVisitor)
	rectangle.Accept(exportVisitor)
	compound.Accept(exportVisitor)
}
