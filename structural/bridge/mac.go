package bridge

import "log"

type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	log.Println("Print request for mac")

	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(printer Printer) {
	m.printer = printer
}
