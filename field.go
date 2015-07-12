package main

import (
	"errors"
	// "fmt"
)

const TOTAL = 9

type Field struct {
	cells []CellType
}

func NewField() *Field {
	field := new(Field)
	field.cells = make([]CellType, 9)
	for i, _ := range field.cells {
		field.cells[i] = Empty
	}

	return field
}

func (field *Field) Set(position int, celltype CellType) error {
	if position > TOTAL-1 {
		return errors.New("index error")
	}
	if field.cells[position] != Empty {
		return errors.New("cant set non empty field")
	}
	field.cells[position] = celltype
	return nil
}

func (field *Field) Clone() *Field {
	_field := NewField()
	copy(_field.cells, field.cells)
	return _field
}

// an immutable version of `Set`, returns error, *Field
func (field *Field) Step(position int, celltype CellType) (error, *Field) {
	_field := field.Clone()
	error := _field.Set(position, celltype)
	return error, _field
}

func (field *Field) Empties() []int {
	empties := make([]int, 0)
	for i, val := range field.cells {
		if val != X && val != O {
			empties = append(empties, i)
		}
	}
	return empties
}

func (field *Field) Get(position int) CellType {
	if position > -1 && position < TOTAL {
		return field.cells[position]
	}
	return Empty
}

func (field *Field) Winner(celltype CellType) bool {
	compbinations := [][]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8},
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8},
		{0, 4, 8}, {2, 4, 6},
	}
	for i := 0; i < len(compbinations); i++ {
		if field.winningCombination(celltype, compbinations[i]) {
			return true
		}
	}
	return false
}

func (field *Field) winningCombination(celltype CellType, strip []int) bool {
	if len(strip) == 3 {
		ret := true
		for i := 1; i < 3; i++ {
			ret = ret && (field.cells[strip[i]] == celltype)
		}
		return ret
	}
	return false
}

func (field *Field) ToString() string {
	var ret string = ""
	for i, value := range field.cells {
		ret += value.ToString()
		if (i+1)%3 == 0 {
			ret += "\n"
		}
	}
	return ret
}
