package table

import (
	"fmt"
	"io"
	"os"

	"robot/internal/direction"
	"robot/internal/point"
)

type Table struct {
	sizeX         uint
	sizeY         uint
	robotPosition *point.Point
	robotFacing   *direction.Direction
	reportOutput  io.Writer
}

// Option is an option that can be passed to `New`
type Option func(*Table)

// WithReportOutput provides an option to specify custom output for the report
func WithReportOutput(out io.Writer) Option {
	return func(t *Table) {
		t.reportOutput = out
	}
}

func New(sizeX, sizeY uint, opts ...Option) *Table {
	tbl := &Table{
		sizeX:        sizeX,
		sizeY:        sizeY,
		reportOutput: os.Stdout,
	}

	for _, opt := range opts {
		opt(tbl)
	}

	return tbl
}

func (t *Table) validatePosition(pos point.Point) error {
	if pos.X < 0 || uint(pos.X) >= t.sizeX {
		return ErrEndingPositionOutOfBounds
	}

	if pos.Y < 0 || uint(pos.Y) >= t.sizeY {
		return ErrEndingPositionOutOfBounds
	}

	return nil
}

func (t *Table) PlaceRobot(pos point.Point, facing direction.Direction) error {
	err := t.validatePosition(pos)
	if err != nil {
		return err
	}

	t.robotPosition = &pos
	t.robotFacing = &facing
	return nil
}

func (t *Table) MoveRobot() (*point.Point, error) {
	if t.robotPosition == nil {
		return nil, ErrUninitializedPlacement
	}

	pos := *t.robotPosition
	pos.X += t.robotFacing.DX()
	pos.Y += t.robotFacing.DY()

	err := t.validatePosition(pos)
	if err != nil {
		return t.robotPosition, err
	}

	t.robotPosition = &pos
	return t.robotPosition, nil
}

func (t *Table) RotateRobot(left bool) (*direction.Direction, error) {
	if t.robotPosition == nil {
		return nil, ErrUninitializedPlacement
	}

	if left {
		t.robotFacing.RotateLeft()
		return t.robotFacing, nil
	}

	t.robotFacing.RotateRight()
	return t.robotFacing, nil
}

func (t *Table) Report() error {
	if t.robotPosition == nil {
		return ErrUninitializedPlacement
	}

	_, err := fmt.Fprintf(t.reportOutput, "Robot position: (%d, %d) facing: %s\n", t.robotPosition.X, t.robotPosition.Y, t.robotFacing)
	return err
}
