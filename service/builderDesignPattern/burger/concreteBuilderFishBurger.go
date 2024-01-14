package burger

// FishBurgerBuilder implements the BurgerBuilder interface for a Fish burger
type FishBurgerBuilder struct {
	burger *FishBurger
}

func NewFishBurgerBuilder() *FishBurgerBuilder {
	return &FishBurgerBuilder{burger: &FishBurger{}}
}

func (b *FishBurgerBuilder) BuildBun() BurgerBuilder {
	b.burger.Bun = "Regular Bun"
	return b
}

func (b *FishBurgerBuilder) BuildPatty() BurgerBuilder {
	b.burger.Patty = "Fish Patty"
	return b
}

func (b *FishBurgerBuilder) AddCondiment(condiment string) BurgerBuilder {
	b.burger.Condiments = append(b.burger.Condiments, condiment)
	return b
}

func (b *FishBurgerBuilder) GetResult() Burger {
	return b.burger
}
