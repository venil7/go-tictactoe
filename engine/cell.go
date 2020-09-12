package engine

type CellType int

const (
	Empty CellType = iota
	X
	O
)

func (celltype CellType) ToString() string {
	var ret string
	switch celltype {
	case X:
		ret = "x"
	case O:
		ret = "o"
	case Empty:
		ret = "_"
	}
	return ret
}

func (celltype CellType) Reverse() CellType {
	var ret CellType
	switch celltype {
	case X:
		ret = O
	case O:
		ret = X
	}
	return ret
}
