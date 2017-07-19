package model

type World struct {
	board   Board
	sausage []Sausage
}

func NewWorld(b Board, s []Sausage) *World {
	return &World{b, s}
}

func (w *World) Board() *Board {
	return &w.board
}

func (w *World) Sausage(i int) *Sausage {
	return &w.sausage[i]
}

func (w *World) Sausages() []Sausage {
	return w.sausage
}
