package burger

// FishBurger represents a fish burger
type FishBurger struct {
	Bun        string
	Patty      string
	Condiments []string
}

func (b *FishBurger) GetBun() string {
	return b.Bun
}

func (b *FishBurger) GetPatty() string {
	return b.Patty
}

func (b *FishBurger) GetCondiments() []string {
	return b.Condiments
}
