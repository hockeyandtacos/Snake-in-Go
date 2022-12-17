package main

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type pos struct {
	x int32
	y int32
}

type Snake struct {
	height int32
	width  int32
	dir    string
	Pos    []pos
}

func (s *Snake) main(game Game, delay *time.Ticker) {
	s.Input()

	<-delay.C

	s.Move(game)
}

func (s *Snake) Input() {
	if rl.IsKeyPressed(rl.KeyUp) && s.dir != "DOWN" {
		s.dir = "UP"
	}
	if rl.IsKeyPressed(rl.KeyDown) && s.dir != "UP" {
		s.dir = "DOWN"
	}
	if rl.IsKeyPressed(rl.KeyRight) && s.dir != "LEFT" {
		s.dir = "RIGHT"
	}
	if rl.IsKeyPressed(rl.KeyLeft) && s.dir != "RIGHT" {
		s.dir = "LEFT"
	}
}

func (s *Snake) addSegment() {

	fmt.Println("addSegment Called")

	lastPos := s.Pos[len(s.Pos)-1]

	fmt.Println("Last Position:", lastPos)

	if s.dir == "UP" {
		s.Pos = append(s.Pos, pos{lastPos.x, lastPos.y + 1})

	}
	if s.dir == "DOWN" {
		s.Pos = append(s.Pos, pos{lastPos.x, lastPos.y - 1})

	}
	if s.dir == "LEFT" {
		s.Pos = append(s.Pos, pos{lastPos.x + 1, lastPos.y})

	}
	if s.dir == "RIGHT" {
		s.Pos = append(s.Pos, pos{lastPos.x - 1, lastPos.y})
		fmt.Println(s.Pos)
	}
}

func (s *Snake) Move(game Game) {
	for i := len(s.Pos) - 1; i > 0; i-- {
		s.Pos[i] = s.Pos[i-1]
	}

	if s.dir == "UP" {
		s.Pos[0].y--
	} else if s.dir == "DOWN" {
		s.Pos[0].y++
	} else if s.dir == "LEFT" {
		s.Pos[0].x--
	} else if s.dir == "RIGHT" {
		s.Pos[0].x++
	}
}

func (s *Snake) head() pos {
	return s.Pos[0]
}
