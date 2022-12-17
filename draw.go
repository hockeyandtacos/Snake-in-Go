package main

import rl "github.com/gen2brain/raylib-go/raylib"

func drawMain(snake Snake, fruit Fruit, game Game) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.White)

	//Snake Head
	rl.DrawRectangle(snake.head().x*game.tile_width, snake.head().y*game.tile_width, snake.width, snake.height, rl.NewColor(0, 165, 230, 255))

	//Snake Body
	for i := 1; i < len(snake.Pos); i++ {
		rl.DrawRectangle(snake.Pos[i].x*game.tile_width, snake.Pos[i].y*game.tile_width, snake.width, snake.height, rl.Blue)
	}

	//Grid Lines
	for i := 0; i < int(game.tile_width); i++ {
		rl.DrawLine(0, int32(i)*game.tile_width, 12*game.tile_width, int32(i)*game.tile_width, rl.Gray)
	}

	// Fruit
	rl.DrawRectangle(fruit.x*game.tile_width, fruit.y*game.tile_height, fruit.width, fruit.height, rl.Red)

	for i := 0; i < int(game.tile_height); i++ {
		rl.DrawLine(int32(i)*game.tile_width, 0, int32(i)*game.tile_width, 12*game.tile_width, rl.Gray)
	}

	//Score
	rl.DrawText("Score: ", 10, 10, 40, rl.Black)

	rl.EndDrawing()
}

func drawTitle(game Game) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.LightGray)

	rl.DrawText("Press Space To Start", 3*game.tile_width, game.window_height/2, 25, rl.Gray)

	rl.EndDrawing()
}
