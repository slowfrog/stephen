package model

// dir values represent the four axis-aligned directions.
type dir struct {
	name string
	dx   int8
	dy   int8
}

// Name returns the name of the direction.
func (d dir) Name() string {
	return d.name
}

// Offset returns the dx, dy offsets of the direction.
func (d dir) Offset() (dx, dy int8) {
	return d.dx, d.dy
}

// The four predefined directions.
var (
	UP    dir = dir{name: "up", dx: 0, dy: -1}
	LEFT  dir = dir{name: "left", dx: -1, dy: 0}
	DOWN  dir = dir{name: "down", dx: 0, dy: 1}
	RIGHT dir = dir{name: "right", dx: 1, dy: 0}
)

// Pos represent the coordinates of a cell on the board.
type Pos struct {
	X int8
	Y int8
}

// Plus offsets the position by the given direction
func (p Pos) Plus(d dir) (q Pos) {
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
