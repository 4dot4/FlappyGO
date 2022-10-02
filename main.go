package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var player = rl.Rectangle{
	X:      400,
	Y:      760,
	Width:  40,
	Height: 40,
}

func main() {
	rl.InitWindow(800, 800, "FlappySpider")
	rl.SetTargetFPS(60)
	const k float32 = 0.1
	var restLength float32 = 400
	var velocity = rl.Vector2{
		X: 400,
		Y: 400,
	}
	var posApoioVec = rl.Vector2{
		X: 400,
		Y: 0,
	}
	for !rl.WindowShouldClose() {
		var posPlayerVec = rl.Vector2{
			X: player.X,
			Y: player.Y,
		}
		player.X = posPlayerVec.X
		player.Y = posPlayerVec.Y
		var force = rl.Vector2Subtract(posPlayerVec, posApoioVec)
		var x = rl.Vector2Distance(posPlayerVec, posApoioVec) - restLength
		force = rl.Vector2Normalize(force)
		force = rl.Vector2Multiply(force, rl.Vector2{X: x, Y: x})
		force = rl.Vector2Multiply(force, rl.Vector2{X: -k, Y: -k})

		// var dt = rl.GetFrameTime()

		velocity = rl.Vector2Add(velocity, force)
		posPlayerVec = rl.Vector2Add(posPlayerVec, velocity)
		velocity = rl.Vector2Multiply(velocity, rl.Vector2{X: 0.99, Y: 0.99})

		rl.BeginDrawing()
		rl.DrawLine(
			player.ToInt32().X+(player.ToInt32().Width/2), 0,
			player.ToInt32().X+(player.ToInt32().Width/2),
			player.ToInt32().Y, rl.Black)
		rl.DrawRectangleRec(player, rl.RayWhite)
		rl.ClearBackground(rl.Blue)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
