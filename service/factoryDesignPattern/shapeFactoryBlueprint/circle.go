package shapeFactoryBlueprint

// 定义具体类型：圆形
type Circle struct {
}

func (c *Circle) Draw() string {
	return "Drawing Circle"
}
