package main

import (
	"time"

	"math/rand"
)

type Fruit struct {
	x      int32
	y      int32
	width  int32
	height int32
	move   bool
}

func (f *Fruit) randomizeLocation() {
	rand.Seed(time.Now().UnixNano())
	f.x = int32(rand.Intn(12))
	f.y = int32(rand.Intn(12))
}

func (f *Fruit) handleMovement(game *Game, snake *Snake) {
	if f.x == snake.head().x && f.y == snake.head().y {
		f.randomizeLocation()
		snake.addSegment()
    game.score++
	}
}
