package engine

type CellType int

const (
	Empty CellType = iota
	X
	O
)

func (celltype CellType) ToString() string {
	switch celltype {
	case X:
		return "x"
	case O:
		return "o"
	}
	return "_"

}

func (celltype CellType) Reverse() CellType {
	switch celltype {
	case X:
		return O
	case O:
		return X
	}
	panic("cant reverse")
}
