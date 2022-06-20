package point

// реализация сеттеров и геттеров

type Point struct {
	x int
	y int
}

func (point *Point) GetX() int {
	return point.x
}

func (point *Point) GetY() int {
	return point.y
}

func (point *Point) SetX(x int) {
	point.x = x
}

func (point *Point) SetY(y int) {
	point.y = y
}
