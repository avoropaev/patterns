package bridge

import "log"

type Linux struct {
	printer Printer
}

func (l *Linux) Print() {
	log.Println("Print request for linux")

	l.printer.PrintFile()
}

func (l *Linux) SetPrinter(printer Printer) {
	l.printer = printer
}
