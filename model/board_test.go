package model

import "testing"

func TestCellName(t *testing.T) {
	var cases = []struct {
		c cell
		s string
	}{
		{EMPTY, "empty"},
		{GROUND, "ground"},
		{GRILL, "grill"},
	}

	for _, tt := range cases {
		if tt.c.Name() != tt.s {
			t.Errorf("EMPTY name should be 'empty', was '%d'", tt.c.Name())
		}
	}
}

func TestCreateGetSetBoard(t *testing.T) {
	var cases = []struct {
		x, y int8
		c cell
	}{
		{0, 0, GROUND},
		{1, 0, EMPTY},
		{2, 0, EMPTY},
		{0, 1, EMPTY},
		{1, 1, EMPTY},
		{2, 1, GRILL},
	}
	b := CreateBoard(3, 2)
	b.Set(0, 0, GROUND).Set(2, 1, GRILL)
	for _, tt := range cases {
		if b.Get(tt.x, tt.y) != tt.c {
			t.Errorf("board.Get(%d, %d) expected %s but got %s",
				tt.x, tt.y, tt.c.Name(), b.Get(tt.x, tt.y).Name())
		}
	}
}

func TestGetOutsideOfBoardReturnsEmpty(t *testing.T) {
	var cases = []Pos {
		Pos{-1, 0},
		Pos{0, -1},
		Pos{-1, -1},
		Pos{1, 2},
		Pos{1, 1},
		Pos{0, 2},
	}
	b := CreateBoard(1, 2)
	for _, tt := range cases {
		if b.Get(tt.X, tt.Y) != EMPTY {
			t.Errorf("Expected empty at %d, %d but was %s",
				tt.X, tt.Y, b.Get(tt.X, tt.Y).Name())
		}
	}
}
