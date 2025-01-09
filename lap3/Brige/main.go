package main

import "fmt"

type Computer interface {
	Print()            // Method Print
	SetPrinter(Printer) 
}

type Printer interface {
	PrintFile() 
}

type Epson struct{}

func (p *Epson) PrintFile() {
	fmt.Println("Đang in bằng máy in EPSON")
}

type Hp struct{}

func (p *Hp) PrintFile() {
	fmt.Println("Đang in bằng máy in HP")
}

type Mac struct {
	printer Printer 
}

func (m *Mac) Print() {
	fmt.Println("Yêu cầu in từ máy Mac")
	m.printer.PrintFile() 
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p 
}

type Windows struct {
	printer Printer 
}

func (w *Windows) Print() {
	fmt.Println("Yêu cầu in từ máy Windows")
	w.printer.PrintFile() 
}

func (w *Windows) SetPrinter(p Printer) {
	w.printer = p 
}

func main() {
	hpPrinter := &Hp{}       
	epsonPrinter := &Epson{} 

	macComputer := &Mac{}
	macComputer.SetPrinter(hpPrinter) 
	macComputer.Print()                
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter) 
	macComputer.Print()                  
	fmt.Println()

	winComputer := &Windows{}
	winComputer.SetPrinter(hpPrinter) 
	winComputer.Print()                
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter) 
	winComputer.Print()                  
	fmt.Println()
}
