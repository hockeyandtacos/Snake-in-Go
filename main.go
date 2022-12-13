package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"time"
)

type Game struct {
	window_height int32
	window_width  int32
	tile_width    int32
	tile_height   int32
	game_over     bool
}

func main() {

	game := Game{600, 600, 50, 50, false}
	snake := Snake{1 * game.tile_width, 5 * game.tile_height, game.tile_height, game.tile_width, "NEUTRAL"}

	delay := time.NewTicker(100 * time.Millisecond)

	rl.InitWindow(game.window_width, game.window_height, "Snake Game")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		for !game.game_over {
			snake.main(game, delay)
			draw(snake)
		}
	}

	rl.CloseWindow()
}
