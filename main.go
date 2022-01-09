package main

import (
	"fmt"
	"github.com/avoropaev/patterns/behavioral/command"
	"github.com/avoropaev/patterns/behavioral/memento"
	"github.com/avoropaev/patterns/behavioral/state"
	"github.com/avoropaev/patterns/creational/abstract_factory"
	"github.com/avoropaev/patterns/creational/builder"
	"github.com/avoropaev/patterns/creational/singleton"
	"github.com/avoropaev/patterns/structural/adapter"
	"github.com/avoropaev/patterns/structural/facade"
	"github.com/avoropaev/patterns/structural/flyweight"
)

func main() {
	fmt.Println("=====")
	fmt.Println("Singleton")
	fmt.Println("=====")
	singleton.Example()

	fmt.Println()
	fmt.Println()
	fmt.Println("=====")
	fmt.Println("Builder")
	fmt.Println("=====")
	builder.Example()

	fmt.Println()
	fmt.Println()
	fmt.Println("=====")
	fmt.Println("Abstract factory")
	fmt.Println("=====")
	abstract_factory.Example()

	fmt.Println()
	fmt.Println()
	fmt.Println("=====")
	fmt.Println("Adapter")
	fmt.Println("=====")
	adapter.Example()

	fmt.Println()
	fmt.Println()
	fmt.Println("=====")
	fmt.Println("Facade")
	fmt.Println("=====")
	facade.Example()

	fmt.Println()
	fmt.Println()
	fmt.Println("=====")
	fmt.Println("Flyweight")
	fmt.Println("=====")
	flyweight.Example()

	fmt.Println()
	fmt.Println()
	fmt.Println("=====")
	fmt.Println("State")
	fmt.Println("=====")
	state.Example()

	fmt.Println()
	fmt.Println()
	fmt.Println("=====")
	fmt.Println("Command")
	fmt.Println("=====")
	command.Example()

	fmt.Println()
	fmt.Println()
	fmt.Println("=====")
	fmt.Println("Memento")
	fmt.Println("=====")
	memento.Example()
}
