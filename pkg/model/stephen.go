package model

import (
	"fmt"
)

type Stephen struct {
	Pos Pos
	Dir Dir
}

func (s Stephen) String() string {
	return fmt.Sprintf("{X:%d Y:%d Dir:%s}",
		s.Pos.X, s.Pos.Y, s.Dir.Name())
}

func (s *Stephen) Move(d Dir) {
	if s.Dir == d || s.Dir == Opposite(d) {
		s.Pos = s.Pos.Plus(d)
	} else { //if s.Dir == TurnClockwise(d) || s.Dir == TurnCounterClockwise(d) {
		s.Dir = d
	}
}
