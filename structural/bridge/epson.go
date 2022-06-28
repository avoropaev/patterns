package bridge

import "log"

type Epson struct{}

func (e *Epson) PrintFile() {
	log.Println("Printing by a EPSON Printer")
}
