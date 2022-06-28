package bridge

import "log"

type HP struct{}

func (h *HP) PrintFile() {
	log.Println("Printing by a HP Printer")
}
