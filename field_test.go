package main

import (
	// "fmt"
	"testing"
)

func _setup(strip []int, celltype CellType) *Field {
	field := NewField()
	for _, i := range strip {
		field.Set(i, celltype)
	}
	return field
}

func _checkWinner(field *Field, celltype CellType, t *testing.T) {
	if field.Winner(celltype) == false {
		t.Fatal("not a winning combination", celltype)
	}
}

// tests:
func TestWinningCombination1(t *testing.T) {
	strip := []int{0, 1, 2}
	field := _setup(strip, X)
	_checkWinner(field, X, t)
}

func TestWinningCombination2(t *testing.T) {
	strip := []int{3, 4, 5}
	field := _setup(strip, X)
	_checkWinner(field, X, t)
}

func TestWinningCombination3(t *testing.T) {
	strip := []int{6, 7, 8}
	field := _setup(strip, X)
	_checkWinner(field, X, t)
}

func TestWinningCombination4(t *testing.T) {
	strip := []int{0, 3, 6}
	field := _setup(strip, X)
	_checkWinner(field, X, t)
}

func TestWinningCombination5(t *testing.T) {
	strip := []int{1, 4, 7}
	field := _setup(strip, X)
	_checkWinner(field, X, t)
}

func TestWinningCombination6(t *testing.T) {
	strip := []int{2, 5, 8}
	field := _setup(strip, X)
	_checkWinner(field, X, t)
}

// setting and getting empty field
func TestSetGetField1(t *testing.T) {
	field := NewField()
	error := field.Set(0, X)
	if error != nil || field.Get(0) != X {
		t.Fatal("set/get failure")
	}
}

// setting already set field
func TestSetGetField2(t *testing.T) {
	field := NewField()
	error := field.Set(4, X)
	error = field.Set(4, O)
	if error == nil || field.Get(4) != X {
		t.Fatal("repeated set/get failure")
	}
}

// not game over when not full
func TestGameOver1(t *testing.T) {
	field := NewField()
	field.Set(4, X)
	if field.GameOver() {
		t.Fatal("game over when not full")
	}
}

// not game over when not full
func TestGameOver2(t *testing.T) {
	field := NewField()
	field.Set(0, X)
	field.Set(1, X)
	field.Set(2, X)
	if !field.GameOver() {
		t.Fatal("winning comb doesnt game over")
	}
}

// game over when full
func TestGameOver3(t *testing.T) {
	field := NewField()
	field.Set(0, X)
	field.Set(1, X)
	field.Set(2, O)
	field.Set(3, O)
	field.Set(4, O)
	field.Set(5, X)
	field.Set(6, O)
	field.Set(7, X)
	field.Set(8, X)
	if !field.GameOver() {
		t.Fatal("winning comb doesnt game over")
	}
}
