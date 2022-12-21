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
  score         int
  prev_score    int
}

func (g *Game) titleinput(game Game) {
	if rl.IsKeyPressed(rl.KeySpace) {
		g.game_over = false
	}
}

func (g *Game) gameOver() {
  g.prev_score = g.score
	g.game_over = true
}

func main() {

	game := Game{600, 600, 50, 50, true, 0, 0}
	snake := Snake{game.tile_height, game.tile_width, "NEUTRAL", []pos{{1, 5}}}
	fruit := Fruit{10, 5, game.tile_width, game.tile_height, false}

	constdelay := time.NewTicker(100 * time.Millisecond)


	rl.InitWindow(game.window_width, game.window_height, "Snake Game")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
    if game.game_over {
      snake = Snake{game.tile_height, game.tile_width, "NEUTRAL", []pos{{1, 5}}} 
      game.score = 0
      drawTitle(game)
			game.titleinput(game)
		} else{
      snake.main(&game, constdelay)
		  fruit.handleMovement(&game, &snake)
		  drawMain(snake, fruit, game)
    }

  }

	rl.CloseWindow()
}
