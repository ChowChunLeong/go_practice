package service

import "go_practice/service/shapeFactoryBlueprint"

// 定义接口
type Shape interface {
	Draw() string
}

// 工厂函数：根据类型参数创建对应的形状
func GetShape(shapeType string) Shape {
	switch shapeType {
	case "circle":
		return &shapeFactoryBlueprint.Circle{}
	case "square":
		return &shapeFactoryBlueprint.Square{}
	default:
		return nil
	}
}
