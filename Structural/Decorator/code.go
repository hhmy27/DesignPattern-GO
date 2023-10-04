package Decorator

// pizza.go: Component interface
type IPizza interface {
	getPrice() int
}

// veggieMania.go: Concrete component
type VeggieMania struct {
}

func (p *VeggieMania) getPrice() int {
	return 15
}

// tomatoTopping.go: Concrete decorator
type TomatoTopping struct {
	pizza IPizza
}

func (t *TomatoTopping) getPrice() int {
	return t.pizza.getPrice() + 7
}

// cheeseTopping.go: Concrete decorator
type CheeseTopping struct {
	pizza IPizza
}

func (c *CheeseTopping) getPrice() int {
	return c.pizza.getPrice() + 10
}
