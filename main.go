package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/kzmrv/go-snake/direction"
	g "github.com/kzmrv/go-snake/geometry"
)

func main() {
	screen := SetupScreen()
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

	maxX, maxY := screen.Size()
	maxDims := g.Point{X: maxX, Y: maxY}
	state := InitState(maxDims)
	SetFood(screen, *state.Food)

	currTick := 500
	ticker := time.NewTicker(time.Duration(currTick) * time.Millisecond)
	go func() {
		for {
			<-ticker.C
			advanceState(&state, screen)
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
				if !direction.IsOpposite(state.Direction, direction.UP) {
					state.Direction = direction.UP
				}
			}
			if ev.Rune() == 'S' || ev.Rune() == 's' {
				if !direction.IsOpposite(state.Direction, direction.DOWN) {
					state.Direction = direction.DOWN
				}
			}
			if ev.Rune() == 'A' || ev.Rune() == 'a' {
				if !direction.IsOpposite(state.Direction, direction.LEFT) {
					state.Direction = direction.LEFT
				}
			}
			if ev.Rune() == 'D' || ev.Rune() == 'd' {
				if !direction.IsOpposite(state.Direction, direction.RIGHT) {
					state.Direction = direction.RIGHT
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
