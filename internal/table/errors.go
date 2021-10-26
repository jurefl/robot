package table

import "errors"

var (
	ErrUninitializedPlacement    error = errors.New("uninitialized placement")
	ErrEndingPositionOutOfBounds error = errors.New("ending position out of bounds")
)
