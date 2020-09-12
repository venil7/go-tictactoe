package main

import Engine "github.com/venil7/gotictactoe/engine"

func main() {
	field := Engine.NewField()
	for !field.GameOver() {
		field.Print()
		field.HumanInput()
		field.CPUInput()
	}
	field.Print()
}
