package Bridge

import (
	"fmt"
	"testing"
)

func TestPrinter(t *testing.T) {
	hpPrinter := &Hp{}
	epsonPrinter := &Epson{}

	macComputer := &Mac{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print() // Printing by a HP Printer
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print() // Printing by a EPSON Printer
	fmt.Println()

	winComputer := &Win{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print() // Printing by a HP Printer
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print() // Printing by a EPSON Printer
	fmt.Println()
}
