package chain_of_responsibility

import "fmt"

func Example() {
	handlers := &ConcreteHandlerA{
		next: &ConcreteHandlerB{
			next: &ConcreteHandlerC{},
		},
	}

	fmt.Println(handlers.SendRequest(2))
}
