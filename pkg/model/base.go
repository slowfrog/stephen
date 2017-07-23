package model

// Dir values represent the four axis-aligned directions.
type Dir byte

const (
	UP    Dir = 0
	RIGHT     = 1
	DOWN      = 2
	LEFT      = 3
)

// Name returns the name of the direction.
func (d Dir) Name() string {
	switch d {
	case UP:
		return "up"
	case RIGHT:
		return "right"
	case DOWN:
		return "down"
	case LEFT:
		return "left"
	default:
		return "WAT?"
	}
}

var dxOffset = [4]int8{0, 1, 0, -1}
var dyOffset = [4]int8{-1, 0, 1, 0}

// Offset returns the dx, dy offsets of the direction.
func (d Dir) Offset() (dx, dy int8) {
	return dxOffset[d], dyOffset[d]
}

func TurnClockwise(d Dir) Dir {
	return (d + 1) % 4
}

func TurnCounterClockwise(d Dir) Dir {
	return (d + 3) % 4
}

func Opposite(d Dir) Dir {
	return (d + 2) % 4
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
