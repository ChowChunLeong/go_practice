package burger

// ChickenBurgerBuilder implements the BurgerBuilder interface for a chicken burger
type ChickenBurgerBuilder struct {
	burger *ChickenBurger
}

func NewChickenBurgerBuilder() *ChickenBurgerBuilder {
	return &ChickenBurgerBuilder{burger: &ChickenBurger{}}
}

func (b *ChickenBurgerBuilder) BuildBun() BurgerBuilder {
	b.burger.Bun = "Regular Bun"
	return b
}

func (b *ChickenBurgerBuilder) BuildPatty() BurgerBuilder {
	b.burger.Patty = "Chicken Patty"
	return b
}

func (b *ChickenBurgerBuilder) AddCondiment(condiment string) BurgerBuilder {
	b.burger.Condiments = append(b.burger.Condiments, condiment)
	return b
}

func (b *ChickenBurgerBuilder) GetResult() Burger {
	return b.burger
}
