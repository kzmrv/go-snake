package main

import (
	"log"
	"sync"
	"time"

	"github.com/gdamore/tcell/v2"
)

type Node struct {
	Prev  *Node
	Point Point
}

type Point struct {
	X int
	Y int
}

func defStyle() tcell.Style {
	return tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
}

func main() {
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
	// Here's how to get the screen size when you need it.
	// xmax, ymax := s.Size()

	// Here's an example of how to inject a keystroke where it will
	// be picked up by the next PollEvent call.  Note that the
	// queue is LIFO, it has a limited length, and PostEvent() can
	// return an error.
	// s.PostEvent(tcell.NewEventKey(tcell.KeyRune, rune('a'), 0))
	directions := make(map[string][2]int)
	directions["Up"] = [2]int{0, -1}
	directions["Down"] = [2]int{0, 1}
	directions["Left"] = [2]int{-1, 0}
	directions["Right"] = [2]int{1, 0}
	direction := directions["Right"]
	mu := sync.Mutex{}
	head := &Node{Prev: nil, Point: Point{X: 1, Y: 1}}
	tail := &Node{Prev: head, Point: Point{X: 0, Y: 1}}
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for {
			<-ticker.C
			mu.Lock()

			newHead := Node{Point: Point{X: head.Point.X + direction[0], Y: head.Point.Y + direction[1]}}
			head.Prev = &newHead
			head = &newHead

			setDef(screen, tail.Point.X, tail.Point.Y)
			tail = tail.Prev

			setSnake(screen, newHead.Point.X, newHead.Point.Y)
			screen.Show()
			mu.Unlock()
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
				direction = directions["Up"]
			}
			if ev.Rune() == 'S' || ev.Rune() == 's' {
				direction = directions["Down"]
			}
			if ev.Rune() == 'A' || ev.Rune() == 'a' {
				direction = directions["Left"]
			}
			if ev.Rune() == 'D' || ev.Rune() == 'd' {
				direction = directions["Right"]
			}
		}
	}
}

func setDef(screen tcell.Screen, x int, y int) {
	screen.SetContent(x, y, ' ', nil, defStyle())
}

func setSnake(screen tcell.Screen, x int, y int) {
	snakeStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorPurple)
	screen.SetContent(x, y, ' ', nil, snakeStyle)
}
