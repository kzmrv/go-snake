package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/kzmrv/go-snake/geometry"
)

func SetupScreen() tcell.Screen {
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

func SetFood(screen tcell.Screen, pt geometry.Point) {
	screen.SetContent(pt.X, pt.Y, ' ', nil, tcell.StyleDefault.Background(tcell.ColorGreen).Foreground(tcell.ColorBlue))
}

func SetDef(screen tcell.Screen, x int, y int) {
	screen.SetContent(x, y, ' ', nil, defStyle())
}

func defStyle() tcell.Style {
	return tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
}

func SetLose(screen tcell.Screen) {
	text := "You lose!"
	style := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorRed)
	x := 1
	y := 1
	for _, rune := range text {
		screen.SetContent(x, y, rune, nil, style)
		x++
	}
}

func SetSnake(screen tcell.Screen, x int, y int) {
	snakeStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorPurple)
	screen.SetContent(x, y, ' ', nil, snakeStyle)
}
