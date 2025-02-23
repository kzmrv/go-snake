package direction

import (
	"github.com/kzmrv/go-snake/geometry"
)

type DirectionId int

const (
	UP    DirectionId = 0
	DOWN  DirectionId = 1
	LEFT  DirectionId = 2
	RIGHT DirectionId = 3
)

func GetDirections() map[DirectionId]geometry.Point {
	mp := make(map[DirectionId]geometry.Point)
	mp[UP] = geometry.Point{X: 0, Y: -1}
	mp[DOWN] = geometry.Point{X: 0, Y: 1}
	mp[LEFT] = geometry.Point{X: -1, Y: 0}
	mp[RIGHT] = geometry.Point{X: 1, Y: 0}
	return mp
}

func IsOpposite(a DirectionId, b DirectionId) bool {
	if a > b {
		a, b = b, a
	}
	return a == UP && b == DOWN || a == LEFT && b == RIGHT
}
