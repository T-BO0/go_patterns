package decorator

// define the interface for the decorator ----->
type ICoffee interface {
	Cost() int
	Description() string
}

// <----- define the interface for the decorator

// define the concrit struct that will be decorated for the decorator ----->
type Coffee struct {
	CoffeeCost        int
	CoffeeDescription string
}

func NewCoffee() ICoffee {
	return &Coffee{
		CoffeeCost:        5,
		CoffeeDescription: "Coffee",
	}
}

func (c *Coffee) Cost() int {
	return c.CoffeeCost
}

func (c *Coffee) Description() string {
	return c.CoffeeDescription
}

// <----- define the concrit struct that will be decorated for the decorator

// define the decorator ----->
type WithMilk struct {
	coffee                ICoffee
	AdditionalCost        int
	AdditionalDescription string
}

func NewWithMilk(coffee ICoffee) ICoffee {
	return &WithMilk{
		coffee:                coffee,
		AdditionalCost:        2,
		AdditionalDescription: "with milk",
	}
}

func (c *WithMilk) Cost() int {
	return c.coffee.Cost() + c.AdditionalCost
}

func (c *WithMilk) Description() string {
	return c.coffee.Description() + " " + c.AdditionalDescription
}

// <----- define the decorator
