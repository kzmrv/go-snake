package main

import (
	"math/rand/v2"

	"github.com/gdamore/tcell/v2"
	"github.com/kzmrv/go-snake/direction"
	g "github.com/kzmrv/go-snake/geometry"
	s "github.com/kzmrv/go-snake/screen"
)

type Node struct {
	Prev  *Node
	Point g.Point
}

type Snake struct {
	Head *Node
	Tail *Node
}

type GameState struct {
	Snake      *Snake
	Food       *g.Point
	Direction  direction.DirectionId
	Directions map[direction.DirectionId]g.Point
	MaxDims    *g.Point
}

func getNewHeadPos(state *GameState) g.Point {
	snake := state.Snake
	directions := state.Directions
	dirVector := directions[state.Direction]
	newX := snake.Head.Point.X + dirVector.X
	newY := snake.Head.Point.Y + dirVector.Y
	if newX > state.MaxDims.X {
		newX = 0
	}
	if newX < 0 {
		newX = state.MaxDims.X
	}
	if newY > state.MaxDims.Y {
		newY = 0
	}
	if newY < 0 {
		newY = state.MaxDims.Y
	}
	return g.Point{X: newX, Y: newY}
}

func advanceState(state *GameState, screen tcell.Screen) {
	snake := state.Snake
	newHeadPos := getNewHeadPos(state)

	if checkCollision(snake.Tail, &newHeadPos) {
		s.SetLose(screen)
		screen.Show()
		return
	}

	newHead := Node{Point: newHeadPos}
	snake.Head.Prev = &newHead
	snake.Head = &newHead

	if newHeadPos.X == state.Food.X && newHeadPos.Y == state.Food.Y {
		state.Food.X = rand.IntN(state.MaxDims.X)
		state.Food.Y = rand.IntN(state.MaxDims.Y)
		s.SetFood(screen, *state.Food)
	} else {
		s.SetDef(screen, state.Snake.Tail.Point.X, state.Snake.Tail.Point.Y)
		snake.Tail = snake.Tail.Prev
	}

	s.SetSnake(screen, newHead.Point.X, newHead.Point.Y)
	screen.Show()
}

func checkCollision(tail *Node, p *g.Point) bool {
	curr := tail
	for curr != nil {
		if g.Equal(&curr.Point, p) {
			return true
		}
		curr = curr.Prev
	}
	return false
}

func InitState(maxDims g.Point) GameState {
	currDirection := direction.RIGHT
	head := &Node{Prev: nil, Point: g.Point{X: 1, Y: 1}}
	tail := &Node{Prev: head, Point: g.Point{X: 0, Y: 1}}
	Snake := &Snake{Head: head, Tail: tail}
	foodX := rand.IntN(maxDims.X)
	foodY := rand.IntN(maxDims.Y)
	Food := &g.Point{X: foodX, Y: foodY}

	return GameState{Food: Food, Snake: Snake, Direction: currDirection, Directions: direction.GetDirections(), MaxDims: &maxDims}
}
