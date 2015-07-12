package main

import (
// "fmt"
)

func main() {
	field := NewField()
	for !field.GameOver() {
		field.Print()
		field.HumanInput()
		field.CPUInput()
	}
	field.Print()
}
