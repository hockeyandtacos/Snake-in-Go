package main

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Snake struct {
	headx  int32
	heady  int32
	height int32
	width  int32
	dir    string
}

func (s *Snake) main(game Game, delay *time.Ticker) {
	s.Input()

	<-delay.C

	s.Move(game)
}

func (s *Snake) Input() {
	if rl.IsKeyPressed(rl.KeyUp) {
		s.dir = "UP"
	}
	if rl.IsKeyPressed(rl.KeyDown) {
		s.dir = "DOWN"
	}
	if rl.IsKeyPressed(rl.KeyRight) {
		s.dir = "RIGHT"
	}
	if rl.IsKeyPressed(rl.KeyLeft) {
		s.dir = "LEFT"
	}
}

func (s *Snake) Move(game Game) {
	if s.dir == "UP" {
		s.heady -= 1 * game.tile_width
	}
	if s.dir == "DOWN" {
		s.heady += 1 * game.tile_width
	}
	if s.dir == "LEFT" {
		s.headx -= 1 * game.tile_width
	}
	if s.dir == "RIGHT" {
		s.headx += 1 * game.tile_width
	}
}
