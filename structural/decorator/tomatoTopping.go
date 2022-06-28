package decorator

type TomatoTopping struct {
	pizza Pizza
}

func (t *TomatoTopping) GetPrice() int {
	return t.pizza.GetPrice() + 5
}
