package decorator

import "log"

func Example() {
	pizza := &VeggieMania{}

	pizzaWithCheese := &CheeseTopping{pizza: pizza}

	pizzaWithCheeseAndTomato := &TomatoTopping{pizza: pizzaWithCheese}

	log.Printf("Цены пиццы с томатами и сыром = %d", pizzaWithCheeseAndTomato.GetPrice())
}
