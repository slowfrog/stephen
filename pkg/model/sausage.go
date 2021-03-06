package model

import (
	"fmt"
)

type alignment struct {
	name string
	dx   int8
	dy   int8
}

var (
	HORIZONTAL alignment = alignment{name: "horizontal", dx: 1, dy: 0}
	VERTICAL   alignment = alignment{name: "vertical", dx: 0, dy: 1}
)

func (a alignment) Name() string {
	return a.name
}

type Sausage struct {
	Alignment alignment
	Pos Pos
	// Baking count: X,Y-bottom, other-bottom, X,Y-top, other top
	Cooked [4]uint8
}

const (
	UNCOOKED = "."
	COOKED   = "x"
	BURNT    = "#"
)

// Returns a 1-char string representing the cooking state of a sausage part
func CookedStr(b uint8) string {
	switch b {
	case 0:
		return UNCOOKED
	case 1:
		return COOKED
	default:
		return BURNT
	}
}

// Creates an uncooked sausage, with the given alignment and position
func CreateSausage(a alignment, x, y int8) Sausage {
	return Sausage{a, Pos{x, y}, [4]uint8{0, 0, 0, 0}}
}

func (s Sausage) GetPos() Pos {
	return s.Pos
}

// Cook a part of the sausage
func (s *Sausage) Cook(which uint8) *Sausage {
	s.Cooked[which] += 1
	return s
}

// Returns a text representation of the sausage
func (s Sausage) String() string {
	return fmt.Sprintf("(%d,%d-%s-[%s%s][%s%s])",
		s.Pos.X, s.Pos.Y, s.Alignment.Name(),
		CookedStr(s.Cooked[0]), CookedStr(s.Cooked[1]),
		CookedStr(s.Cooked[2]), CookedStr(s.Cooked[3]))
}
