package main

import (
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

func (s *Snake) main(game *Game, delay *time.Ticker) {
	s.Input()

	<-delay.C

	s.Move(game)
  s.collision(game)
  s.outOfBounds(game)
}

func (s *Snake) Input() {
  var moved bool = false

	if rl.IsKeyPressed(rl.KeyUp) && s.dir != "DOWN" && !moved{
		s.dir = "UP"
    moved = true
  } else if rl.IsKeyPressed(rl.KeyDown) && s.dir != "UP" && !moved{
		s.dir = "DOWN"
    moved = true
	} else if rl.IsKeyPressed(rl.KeyRight) && s.dir != "LEFT" && !moved {
		s.dir = "RIGHT"
    moved = true
	} else if rl.IsKeyPressed(rl.KeyLeft) && s.dir != "RIGHT" && !moved{
		s.dir = "LEFT"
    moved = true
	}
}

func (s *Snake) addSegment() {

	lastPos := s.Pos[len(s.Pos)-1]

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
	}
}

func (s *Snake) Move(game *Game) {
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

func (s *Snake) collision(game *Game) {
  var touch bool

  for i := 1; i  < len(s.Pos); i ++ {
    if s.Pos[i].x == s.head().x && s.Pos[i].y == s.head().y {
      touch = true
    }
  }

  if touch {
    game.gameOver()
  }
}

func (s *Snake) outOfBounds(game *Game) {
  if s.head().x < 0 || s.head().x > 11 || s.head().y < 0 || s.head().y > 11 {
    game.gameOver()
  }
}

