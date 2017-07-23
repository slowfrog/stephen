package model

type World struct {
	board   Board
	sausage []Sausage
	stephen Stephen
}

func NewWorld(b Board, s []Sausage, st Stephen) *World {
	return &World{b, s, st}
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

func (w *World) Stephen() *Stephen {
	return &w.stephen
}

func (w *World) MoveStephen(d Dir) {
	w.stephen.Move(d)
}
