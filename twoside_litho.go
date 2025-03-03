package main

import (
	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	gpio "github.com/syferever/twoside-litho-gui/gpioctl"
)

func main() {

	rl.InitWindow(1920, 1080, "Two Side Litho")
	var left_out_led int32
	var left_in_led int32
	var right_out_led int32
	var right_in_led int32
	left_pic := rl.LoadImage("./assets/left.jpg")
	rl.ImageResize(left_pic, 930, 850)
	left_texture := rl.LoadTextureFromImage(left_pic)
	right_pic := rl.LoadImage("./assets/right.jpg")
	rl.ImageResize(right_pic, 930, 850)
	right_texture := rl.LoadTextureFromImage(right_pic)

	gpio.HostInit()
	left_out_pin := gpio.NewPin("GPIO17")
	left_out_pin.SetPinMode("OUT")

	for !rl.WindowShouldClose() {

		switch {
		case left_out_led == 1 && left_out_pin.Level != gpio.HIGH:
			left_out_pin.SetPinLevel(gpio.HIGH)
		case left_out_led == 0 && left_out_pin.Level != gpio.LOW:
			left_out_pin.SetPinLevel(gpio.LOW)
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		rl.DrawRectangleLines(10, 30, 220, 100, rl.Black)
		rl.DrawText("Left", 100, 40, 20, rl.Black)
		rl.DrawText("In", 60, 70, 20, rl.Black)
		rl.DrawText("Out", 160, 70, 20, rl.Black)
		rl.DrawRectangleLines(240, 30, 220, 100, rl.Black)
		rl.DrawText("Right", 330, 40, 20, rl.Black)
		rl.DrawText("In", 290, 70, 20, rl.Black)
		rl.DrawText("Out", 380, 70, 20, rl.Black)
		left_in_led = rg.ToggleSlider(rl.NewRectangle(20, 95, 90, 25), "Off;On", left_in_led)
		left_out_led = rg.ToggleSlider(rl.NewRectangle(130, 95, 90, 25), "Off;On", left_out_led)
		right_in_led = rg.ToggleSlider(rl.NewRectangle(250, 95, 90, 25), "Off;On", right_in_led)
		right_out_led = rg.ToggleSlider(rl.NewRectangle(350, 95, 90, 25), "Off;On", right_out_led)

		rl.DrawTexture(left_texture, 20, 150, rl.White)
		rl.DrawTexture(right_texture, 970, 150, rl.White)

		rl.EndDrawing()
	}
}
