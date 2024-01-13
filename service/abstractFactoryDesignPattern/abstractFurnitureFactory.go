package abstractfactorydesignpattern

type FurnitureFactory interface {
	CreateChair() Chair
	CreateTable() Table
}

// Abstract Product Interfaces
type Chair interface {
	SitOn() string
}

type Table interface {
	PutOn() string
}
