package main

import (
	"fmt"

	"github.com/avoropaev/patterns/behavioral/chain_of_responsibility"
	"github.com/avoropaev/patterns/behavioral/command"
	"github.com/avoropaev/patterns/behavioral/memento"
	"github.com/avoropaev/patterns/behavioral/state"
	"github.com/avoropaev/patterns/behavioral/strategy"
	"github.com/avoropaev/patterns/behavioral/template_method"
	"github.com/avoropaev/patterns/creational/abstract_factory"
	"github.com/avoropaev/patterns/creational/builder"
	"github.com/avoropaev/patterns/creational/factory_method"
	"github.com/avoropaev/patterns/creational/prototype"
	"github.com/avoropaev/patterns/creational/singleton"
	"github.com/avoropaev/patterns/structural/adapter"
	"github.com/avoropaev/patterns/structural/bridge"
	"github.com/avoropaev/patterns/structural/composite"
	"github.com/avoropaev/patterns/structural/decorator"
	"github.com/avoropaev/patterns/structural/facade"
	"github.com/avoropaev/patterns/structural/flyweight"
	"github.com/avoropaev/patterns/structural/proxy"
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
	fmt.Println("Factory method")
	fmt.Println("=====")
	factory_method.Example()

	fmt.Println()
	fmt.Println()
	fmt.Println("=====")
	fmt.Println("Prototype")
	fmt.Println("=====")
	prototype.Example()

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
	fmt.Println("Proxy")
	fmt.Println("=====")
	proxy.Example()

	fmt.Println()
	fmt.Println()
	fmt.Println("=====")
	fmt.Println("Bridge")
	fmt.Println("=====")
	bridge.Example()

	fmt.Println()
	fmt.Println()
	fmt.Println("=====")
	fmt.Println("Composite")
	fmt.Println("=====")
	composite.Example()

	fmt.Println()
	fmt.Println()
	fmt.Println("=====")
	fmt.Println("Decorator")
	fmt.Println("=====")
	decorator.Example()

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

	fmt.Println()
	fmt.Println()
	fmt.Println("=====")
	fmt.Println("Strategy")
	fmt.Println("=====")
	strategy.Example()

	fmt.Println()
	fmt.Println()
	fmt.Println("=====")
	fmt.Println("Chain of Responsibility")
	fmt.Println("=====")
	chain_of_responsibility.Example()

	fmt.Println()
	fmt.Println()
	fmt.Println("=====")
	fmt.Println("Template method")
	fmt.Println("=====")
	template_method.Example()
}
