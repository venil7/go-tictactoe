package engine

import (
	"errors"
	"fmt"
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
	if position < 0 || position >= TOTAL {
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

func (field *Field) Eval(depth int) int {
	if field.Winner(O) {
		return +10 - depth
	}
	if field.Winner(X) {
		return depth - 10
	}

	return 0
}

// an immutable version of `Set`, returns error, *Field
func (field *Field) Step(position int, celltype CellType) (*Field, error) {
	_field := field.Clone()
	error := _field.Set(position, celltype)
	return _field, error
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
		for i := 0; i < 3; i++ {
			ret = ret && (field.cells[strip[i]] == celltype)
		}
		return ret
	}
	return false
}

func (field *Field) GameOver() bool {
	return field.Winner(X) || field.Winner(O) || len(field.Empties()) == 0
}

func (field *Field) Print() {
	fmt.Print("\n", field.ToString(), "\n")
}

func (field *Field) HumanInput() {
	for {
		var i int
		_, error := fmt.Scanf("%d", &i)
		if error != nil {
			continue
		}
		error = field.Set(i, X)
		if error != nil {
			fmt.Errorf("%s", error)
			continue
		} else {
			return
		}
	}
}

func maxindex(arr []int) int {
	idx, val := 0, arr[0]
	for index, value := range arr {
		if value > val {
			idx, val = index, value
		}
	}
	return idx
}

func minindex(arr []int) int {
	idx, val := 0, arr[0]
	for index, value := range arr {
		if value < val {
			idx, val = index, value
		}
	}
	return idx
}

func (field *Field) CPUInput() {
	_, move := Minimax(field, X, -1, 0)
	field.Set(move, O)
}

func Minimax(field *Field, celltype CellType, pos int, depth int) (int, int) /*score, position*/ {
	scores := make([]int, 0)
	moves := make([]int, 0)

	if field.GameOver() {
		return field.Eval(depth), pos
	}

	for _, move := range field.Empties() {
		_field, error := field.Step(move, celltype)
		if error == nil {
			score, _ := Minimax(_field, celltype.Reverse(), move, depth+1)
			scores = append(scores, score)
			moves = append(moves, move)
		} else {
			fmt.Println(error)
		}
	}

	var idx int
	switch celltype {
	case X:
		idx = minindex(scores)
	case O:
		idx = maxindex(scores)
	}

	return scores[idx], moves[idx]

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
