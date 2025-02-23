package geometry

type Point struct {
	X int
	Y int
}

func Equal(p1 *Point, p2 *Point) bool {
	return p1.X == p2.X && p1.Y == p2.Y
}
