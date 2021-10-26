package direction

// Direction defines by how much the object would advance alongise X and Y axis
type Direction struct {
	dX int
	dY int
}

var (
	East  = Direction{dX: 1, dY: 0}
	North = Direction{dX: 0, dY: 1}
	West  = Direction{dX: -1, dY: 0}
	South = Direction{dX: 0, dY: -1}

	// All directions in counterclockwise order
	ccwDirections = [...]Direction{East, North, West, South}
)

// String returns a string representation of Direction
func (d Direction) String() string {
	switch d {
	case East:
		return "EAST"
	case North:
		return "NORTH"
	case West:
		return "WEST"
	default:
		return "SOUTH"
	}
}

// RotateLeft rotates direction by 90 degress counterclockwise
func (d *Direction) RotateLeft() {
	for i, dr := range ccwDirections {
		if dr == *d {
			leftIdx := (i + 1) % len(ccwDirections)
			*d = ccwDirections[leftIdx]
			break
		}
	}
}

// RotateRight rotates direction by 90 degress clockwise
func (d *Direction) RotateRight() {
	for i, dr := range ccwDirections {
		if dr == *d {
			ccwDirectionsLen := len(ccwDirections)
			rightIdx := (i + ccwDirectionsLen - 1) % ccwDirectionsLen
			*d = ccwDirections[rightIdx]
			break
		}
	}
}

func (d Direction) DX() int {
	return d.dX
}

func (d Direction) DY() int {
	return d.dY
}
