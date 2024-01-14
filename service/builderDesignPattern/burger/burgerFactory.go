package burger

// BurgerFactory is a factory function that creates the appropriate builder based on the input
func BurgerFactory(burgerType string) BurgerBuilder {
	switch burgerType {
	case "chicken":
		return NewChickenBurgerBuilder()
	case "fish":
		return NewFishBurgerBuilder()
	default:
		return nil
	}
}
