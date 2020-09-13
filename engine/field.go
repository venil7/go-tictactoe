package engine

import (
	"errors"
	"fmt"
	"sort"
)

const TOTAL = 9

type Field struct {
	cells []CellType
}

type Eval struct {
	score    int
	position int
}

func NewField() *Field {
	field := new(Field)
	field.cells = make([]CellType, 9)

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
	combinations := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}
	ch := make(chan bool /*, len(combinations)*/)
	for _, combination := range combinations {
		go func(celltype CellType, comb []int, result chan<- bool) {
			ch <- field.winningCombination(celltype, comb)
		}(celltype, combination, ch)
	}
	for i := 0; i < len(combinations); i++ {
		select {
		case ret := <-ch:
			if ret {
				return ret
			}
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

func (field *Field) CPUInput() {
	// eval := Minimax(field, X, -1, 0) // no goroutines version
	eval := minimax(field)
	field.Set(eval.position, O)
}

func minimax(field *Field) Eval {
	empties := field.Empties()
	len := len(empties)
	ch := make(chan Eval, len)
	for _, possiblePosition := range empties {
		go func(field *Field, possiblePosition int, result chan<- Eval) {
			fieldCopy, _ := field.Step(possiblePosition, O /*cpu*/)
			result <- Minimax(fieldCopy, X /*human*/, possiblePosition, 1)
		}(field, possiblePosition, ch)
	}

	var eval Eval = Eval{score: -1000}

	for i := 0; i < len; i++ {
		select {
		case ev := <-ch:
			{
				if ev.score > eval.score {
					eval = ev
				}
			}
		}
	}

	return eval

}

func Minimax(field *Field, celltype CellType, pos int, depth int) Eval {
	evals := make([]Eval, 0)

	if field.GameOver() {
		return Eval{score: field.Eval(depth), position: pos}
	}

	for _, possiblePosition := range field.Empties() {
		fieldCopy, error := field.Step(possiblePosition, celltype)
		if error == nil {
			eval := Minimax(fieldCopy, celltype.Reverse(), possiblePosition, depth+1)
			evals = append(evals, eval)
		} else {
			fmt.Errorf("%s", error)
		}
	}

	sort.Slice(evals, func(i, j int) bool { return evals[i].score < evals[j].score })

	if celltype == X {
		return Eval{score: evals[0].score, position: pos}
	}
	return Eval{score: evals[len(evals)-1].score, position: pos}
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
