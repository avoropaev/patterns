package decorator

type CheeseTopping struct {
	pizza Pizza
}

func (c *CheeseTopping) GetPrice() int {
	return c.pizza.GetPrice() + 7
}
