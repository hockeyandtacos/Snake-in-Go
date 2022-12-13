package main

import rl "github.com/gen2brain/raylib-go/raylib"

func draw(snake Snake) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	rl.DrawRectangle(snake.headx, snake.heady, snake.width, snake.height, rl.Green) // Snake

	rl.EndDrawing()
}
