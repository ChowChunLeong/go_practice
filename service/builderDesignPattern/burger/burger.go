package burger

// Burger represents the common interface for all burgers
type Burger interface {
	GetBun() string
	GetPatty() string
	GetCondiments() []string
}

// BurgerBuilder is the common interface for burger builders
type BurgerBuilder interface {
	BuildBun() BurgerBuilder
	BuildPatty() BurgerBuilder
	AddCondiment(condiment string) BurgerBuilder
	GetResult() Burger
}
