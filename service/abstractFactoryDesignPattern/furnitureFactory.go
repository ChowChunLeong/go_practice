package abstractfactorydesignpattern

func CreateFurniture(variant string, productType string) (Chair, Table) {
	var factory FurnitureFactory

	switch variant {
	case "Modern":
		factory = ModernFurnitureFactory{}
	case "Victorian":
		factory = VictorianFurnitureFactory{}
	default:
		// Handle unknown style
		return nil, nil
	}

	switch productType {
	case "Chair":
		return factory.CreateChair(), nil
	case "Table":
		return nil, factory.CreateTable()
	default:
		// Handle unknown product type
		return nil, nil
	}
}
