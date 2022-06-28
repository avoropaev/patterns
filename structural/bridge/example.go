package bridge

import "fmt"

func Example() {
	hp := HP{}
	epson := Epson{}

	mac := Mac{}
	mac.SetPrinter(&hp)
	mac.Print()
	fmt.Println()
	mac.SetPrinter(&epson)
	mac.Print()
	fmt.Println()
	fmt.Println()
	linux := Linux{}
	linux.SetPrinter(&hp)
	linux.Print()
	fmt.Println()
	linux.SetPrinter(&epson)
	linux.Print()
}
