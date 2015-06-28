package main

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

func (field *Field) Winner(celltype CellType) bool {
	compbinations := [][]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8},
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8},
		{0, 4, 8}, {2, 4, 6},
	}

	for i := 0; i < len(compbinations); i++ {
		if field.WinningCombination(celltype, compbinations[i]) {
			return true
		}
	}

	return false
}

func (field *Field) WinningCombination(celltype CellType, strip []int) bool {
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
