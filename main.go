package main

import (
	"log"
	"math/rand/v2"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/kzmrv/go-snake/direction"
	g "github.com/kzmrv/go-snake/geometry"
)

type Node struct {
	Prev  *Node
	Point g.Point
}

func defStyle() tcell.Style {
	return tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
}

func main() {
	screen := setupScreen()
	quit := func() {
		// You have to catch panics in a defer, clean up, and
		// re-raise them - otherwise your application can
		// die without leaving any diagnostic trace.
		maybePanic := recover()
		screen.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()
	directions := direction.GetDirections()
	currDirection := direction.RIGHT
	maxX, maxY := screen.Size()

	head := &Node{Prev: nil, Point: g.Point{X: 1, Y: 1}}
	tail := &Node{Prev: head, Point: g.Point{X: 0, Y: 1}}

	foodX := rand.IntN(maxX)
	foodY := rand.IntN(maxY)
	setFood(screen, foodX, foodY)

	currTick := 500
	ticker := time.NewTicker(time.Duration(currTick) * time.Millisecond)
	go func() {
		for {
			<-ticker.C

			dirVector := directions[currDirection]
			newX := head.Point.X + dirVector.X
			newY := head.Point.Y + dirVector.Y
			if newX > maxX {
				newX = 0
			}
			if newX < 0 {
				newX = maxX
			}
			if newY > maxY {
				newY = 0
			}
			if newY < 0 {
				newY = maxY
			}

			if checkCollision(tail, newX, newY) {
				setLose(screen)
				return
			}

			newHead := Node{Point: g.Point{X: newX, Y: newY}}
			head.Prev = &newHead
			head = &newHead

			if newX == foodX && newY == foodY {
				foodX = rand.IntN(maxX)
				foodY = rand.IntN(maxY)
				setFood(screen, foodX, foodY)
			} else {
				setDef(screen, tail.Point.X, tail.Point.Y)
				tail = tail.Prev
			}

			setSnake(screen, newHead.Point.X, newHead.Point.Y)
			screen.Show()
		}
	}()
	// Event loop
	for {
		// Update screen
		screen.Show()
		// Poll event
		ev := screen.PollEvent()

		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			screen.Sync()
		case *tcell.EventKey:
			if ev.Rune() == 'C' || ev.Rune() == 'c' {
				return
			}
			if ev.Rune() == 'W' || ev.Rune() == 'w' {
				if !direction.IsOpposite(currDirection, direction.UP) {
					currDirection = direction.UP
				}
			}
			if ev.Rune() == 'S' || ev.Rune() == 's' {
				if !direction.IsOpposite(currDirection, direction.DOWN) {
					currDirection = direction.DOWN
				}
			}
			if ev.Rune() == 'A' || ev.Rune() == 'a' {
				if !direction.IsOpposite(currDirection, direction.LEFT) {
					currDirection = direction.LEFT
				}
			}
			if ev.Rune() == 'D' || ev.Rune() == 'd' {
				if !direction.IsOpposite(currDirection, direction.RIGHT) {
					currDirection = direction.RIGHT
				}
			}
			if ev.Rune() == '[' {
				currTick /= 2
				ticker.Reset(time.Duration(currTick) * time.Millisecond)
			}
			if ev.Rune() == ']' {
				currTick = currTick*2 + 1
				ticker.Reset(time.Duration(currTick) * time.Millisecond)
			}
		}
	}
}

func checkCollision(head *Node, x, y int) bool {
	curr := head
	for curr != nil {
		if curr.Point.X == x && curr.Point.Y == y {
			return true
		}
		curr = curr.Prev
	}
	return false
}

func setDef(screen tcell.Screen, x int, y int) {
	screen.SetContent(x, y, ' ', nil, defStyle())
}

func setLose(screen tcell.Screen) {
	text := "You lose!"
	style := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorRed)
	x := 1
	y := 1
	for _, rune := range text {
		screen.SetContent(x, y, rune, nil, style)
		x++
	}
}

func setSnake(screen tcell.Screen, x int, y int) {
	snakeStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorPurple)
	screen.SetContent(x, y, ' ', nil, snakeStyle)
}

func setFood(screen tcell.Screen, x int, y int) {
	screen.SetContent(x, y, ' ', nil, tcell.StyleDefault.Background(tcell.ColorGreen).Foreground(tcell.ColorBlue))
}

func setupScreen() tcell.Screen {
	defStyle := defStyle()

	// Initialize screen
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	screen.SetStyle(defStyle)
	screen.Clear()
	return screen
}
