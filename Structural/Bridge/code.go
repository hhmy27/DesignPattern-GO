package Bridge

import "fmt"

// printer.go
type Printer interface {
	PrintFile()
}

// computer.go
type Computer interface {
	Print()
	SetPrinter(p Printer)
}

// mac.go
type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

// win.go
type Win struct {
	printer Printer
}

func (w *Win) Print() {
	w.printer.PrintFile()
}

func (w *Win) SetPrinter(p Printer) {
	w.printer = p
}

// epson.go
type Epson struct {
}

func (p *Epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}

// hp.go
type Hp struct {
}

func (p *Hp) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}
