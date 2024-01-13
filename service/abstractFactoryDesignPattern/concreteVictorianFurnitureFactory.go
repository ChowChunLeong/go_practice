package abstractfactorydesignpattern

type VictorianFurnitureFactory struct{}

func (v VictorianFurnitureFactory) CreateChair() Chair {
	return &VictorianChair{}
}

func (v VictorianFurnitureFactory) CreateTable() Table {
	return &VictorianTable{}
}

// Concrete Product 1: Victorian Chair
type VictorianChair struct{}

func (v VictorianChair) SitOn() string {
	return "Sitting on a Victorian chair"
}

// Concrete Product 2: Victorian Table
type VictorianTable struct{}

func (v VictorianTable) PutOn() string {
	return "Putting things on a Victorian table"
}
