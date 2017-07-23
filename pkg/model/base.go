package model

// Dir values represent the four axis-aligned directions.
type Dir struct {
	name string
	dx   int8
	dy   int8
}

// Name returns the name of the direction.
func (d Dir) Name() string {
	return d.name
}

// Offset returns the dx, dy offsets of the direction.
func (d Dir) Offset() (dx, dy int8) {
	return d.dx, d.dy
}

// The four predefined directions.
var (
	UP    Dir = Dir{name: "up", dx: 0, dy: -1}
	LEFT  Dir = Dir{name: "left", dx: -1, dy: 0}
	DOWN  Dir = Dir{name: "down", dx: 0, dy: 1}
	RIGHT Dir = Dir{name: "right", dx: 1, dy: 0}
)

func TurnClockwise(d Dir) Dir {
	switch d {
	case UP:
		return RIGHT
	case RIGHT:
		return DOWN
	case DOWN:
		return LEFT
	case LEFT:
		return UP
	}
	return UP
}

func TurnCounterClockwise(d Dir) Dir {
	switch d {
	case UP:
		return LEFT
	case LEFT:
		return DOWN
	case DOWN:
		return RIGHT
	case RIGHT:
		return UP
	}
	return UP
}

func Opposite(d Dir) Dir {
	switch d {
	case UP:
		return DOWN
	case LEFT:
		return RIGHT
	case DOWN:
		return UP
	case RIGHT:
		return LEFT
	}
	return LEFT
}

// Pos represent the coordinates of a cell on the board.
type Pos struct {
	X int8
	Y int8
}

// Plus offsets the position by the given direction
func (p Pos) Plus(d Dir) (q Pos) {
	dx, dy := d.Offset()
	q.X = p.X + dx
	q.Y = p.Y + dy
	return
}

// moves defines possible character actions: rotations or change of position
type move struct {
	name string
}

var (
	FORWARD    move = move{name: "forward"}
	BACKWARD   move = move{name: "backward"}
	TURN_LEFT  move = move{name: "turn left"}
	TURN_RIGHT move = move{name: "turn right"}
)
