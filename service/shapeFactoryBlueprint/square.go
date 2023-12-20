package shapeFactoryBlueprint

// 定义具体类型：正方形
type Square struct{}

func (s *Square) Draw() string {
	return "Drawing Square"
}
