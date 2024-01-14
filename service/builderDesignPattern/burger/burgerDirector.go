package burger

// BurgerDirector constructs burgers using a builder
type BurgerDirector struct {
	builder BurgerBuilder
}

func NewBurgerDirector(builder BurgerBuilder) *BurgerDirector {
	return &BurgerDirector{builder: builder}
}

func (d *BurgerDirector) Construct() Burger {
	return d.builder.BuildBun().BuildPatty().AddCondiment("Ketchup").AddCondiment("Lettuce").GetResult()
}
