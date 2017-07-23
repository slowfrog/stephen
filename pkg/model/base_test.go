package model

import "testing"

func TestPredefinedDirs(t *testing.T) {
	var cases = []struct {
		d    Dir
		name string
		dx   int8
		dy   int8
	}{
		{UP, "up", 0, -1},
		{RIGHT, "right", 1, 0},
		{DOWN, "down", 0, 1},
		{LEFT, "left", -1, 0},
	}

	for _, tt := range cases {
		if tt.d.Name() != tt.name {
			t.Errorf("Wrong name, expected %s but was %s\n", tt.name, tt.d.Name())
		}
		dx, dy := tt.d.Offset()
		if dx != tt.dx || dy != tt.dy {
			t.Errorf("%s should go %d,%d, was %d,%d\n", tt.name, tt.dx, tt.dy, dx, dy)
		}
	}
}

func TestPosPlusDir(t *testing.T) {
	p := Pos{X: 5, Y: 7}
	var cases = []struct {
		d   Dir
		exp Pos
	}{
		{UP, Pos{5, 6}},
		{RIGHT, Pos{6, 7}},
		{DOWN, Pos{5, 8}},
		{LEFT, Pos{4, 7}},
	}

	for _, tt := range cases {
		a := p.Plus(tt.d)
		if a.X != tt.exp.X || a.Y != tt.exp.Y {
			t.Errorf("Wrong pos+%s, expected %d,%d, was %d,%d",
				tt.d.Name(), tt.exp.X, tt.exp.Y, a.X, a.Y)
		}
	}
}

func TestTurnCW(t *testing.T) {
	var cases = []struct {
		from Dir
		to   Dir
	}{
		{UP, RIGHT},
		{RIGHT, DOWN},
		{DOWN, LEFT},
		{LEFT, UP},
	}

	for _, tt := range cases {
		to := TurnClockwise(tt.from)
		if to != tt.to {
			t.Errorf("Wrong CW turn from %s, expected %s, got %s",
				tt.from.Name(), tt.to.Name(), to.Name())
		}
		from := TurnCounterClockwise(tt.to)
		if from != tt.from {
			t.Errorf("Wrong CCW turn from %s, expected %s, got %s",
				tt.to.Name(), tt.from.Name(), from.Name())
		}
	}
}
