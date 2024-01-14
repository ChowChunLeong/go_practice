package burger

// ChickenBurger represents a chicken burger
type ChickenBurger struct {
	Bun        string
	Patty      string
	Condiments []string
}

func (b *ChickenBurger) GetBun() string {
	return b.Bun
}

func (b *ChickenBurger) GetPatty() string {
	return b.Patty
}

func (b *ChickenBurger) GetCondiments() []string {
	return b.Condiments
}
