package abstractfactorydesignpattern

type ModernFurnitureFactory struct{}

func (m ModernFurnitureFactory) CreateChair() Chair {
	return &ModernChair{}
}

func (m ModernFurnitureFactory) CreateTable() Table {
	return &ModernTable{}
}

// Concrete Product 1: Modern Chair
type ModernChair struct{}

func (m ModernChair) SitOn() string {
	return "Sitting on a modern chair"
}

// Concrete Product 2: Modern Table
type ModernTable struct{}

func (m ModernTable) PutOn() string {
	return "Putting things on a modern table"
}
