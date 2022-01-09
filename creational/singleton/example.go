package singleton

import "fmt"

func Example() {
	for i := 0; i < 7; i++ {
		go getInstance()
	}

	_, err := fmt.Scanln()
	if err != nil {
		return
	}
}
