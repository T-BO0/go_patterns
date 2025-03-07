package decorator

// ICoffee is an interface that will be implemented by all the concrete coffee types
type ICoffee interface {
	Cost() int
	Description() string
}

// Coffee is a struct that will be used as a concrete coffee type
type Coffee struct {
	CoffeeCost        int
	CoffeeDescription string
}

// WithMilk is a struct that will be used as a concrete coffee type with milk. It will implement the ICoffee interface
// WithMilk will have an additional cost and description
// WithMilk will have a reference to the ICoffee interface
type WithMilk struct {
	coffee                ICoffee
	AdditionalCost        int
	AdditionalDescription string
}

// Cost is a method that will return the cost of the coffee with milk (additional cost + coffee cost)
func (c *WithMilk) Cost() int {
	return c.coffee.Cost() + c.AdditionalCost
}

// Description is a method that will return the description of the coffee with milk (additional description + coffee description)
func (c *WithMilk) Description() string {
	return c.coffee.Description() + " " + c.AdditionalDescription
}
