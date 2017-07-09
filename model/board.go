package model

// cell values are the fixed background types of cells
type cell uint8

// All cell values
const (
	EMPTY cell = iota
	GROUND
	GRILL
)

func (c cell) Name() string {
	switch c {
	case EMPTY:
		return "empty"
	case GROUND:
		return "ground"
	case GRILL:
		return "grill"
	default:
		return "WAT?"
	}
}

// Board is the whole static world
type Board struct {
	width  uint8
	height uint8
	cell   []cell
}

// Create a board of a given size, filled with empties
func CreateBoard(width, height uint8) (b Board) {
	b.width = width
	b.height = height
	b.cell = make([]cell, b.width * b.height)
	for i := range b.cell {
		b.cell[i] = EMPTY
	}
	return
}

func (b *Board) Width() uint8 {
	return b.width
}

func (b *Board) Height() uint8 {
	return b.height
}

func (b *Board) Size() (w, h uint8) {
	w = b.width
	h = b.height
	return
}

func (b *Board) index(x, y uint8) uint8 {
	return b.width * y + x
}

// Sets a value at a given position
func (b *Board) Set(x, y uint8, c cell) *Board {
	i := b.index(x, y)
	b.cell[i] = c
	return b
}

// Gets the value at a given position
func (b *Board) Get(x, y int8) cell {
	if x < 0 || uint8(x) >= b.width || y < 0 || uint8(y) >= b.height {
		return EMPTY
	}
	return b.cell[b.index(uint8(x), uint8(y))]
}
