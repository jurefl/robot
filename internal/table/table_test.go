package table_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	"robot/internal/direction"
	"robot/internal/point"
	"robot/internal/table"
)

func TestPlaceRobot(t *testing.T) {
	t.Parallel()

	tests := [...]struct {
		name     string
		tbl      *table.Table
		pos      point.Point
		facing   direction.Direction
		expected error
	}{
		{
			name:     "should successfully place a robot on the board",
			tbl:      table.New(5, 5),
			pos:      point.Point{X: 1, Y: 1},
			facing:   direction.West,
			expected: nil,
		},
		{
			name:     "should fail to place a robot when X position is out of bounds",
			tbl:      table.New(5, 5),
			pos:      point.Point{X: 10, Y: 1},
			facing:   direction.West,
			expected: table.ErrEndingPositionOutOfBounds,
		},
		{
			name:     "should fail to place a robot when Y position is out of bounds",
			tbl:      table.New(5, 5),
			pos:      point.Point{X: 0, Y: -1},
			facing:   direction.West,
			expected: table.ErrEndingPositionOutOfBounds,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := tt.tbl.PlaceRobot(tt.pos, tt.facing)
			require.Equal(t, tt.expected, actual)
		})
	}
}

func TestMoveRobot(t *testing.T) {
	t.Parallel()

	tests := [...]struct {
		name        string
		tbl         func() *table.Table
		expectedPos *point.Point
		expectedErr error
	}{
		{
			name: "should ignore move command when placement was not done",
			tbl: func() *table.Table {
				tbl := table.New(5, 5)
				return tbl
			},
			expectedPos: nil,
			expectedErr: table.ErrUninitializedPlacement,
		},
		{
			name: "should successfully move a robot",
			tbl: func() *table.Table {
				tbl := table.New(5, 5)
				tbl.PlaceRobot(point.Point{X: 1, Y: 1}, direction.North)
				return tbl
			},
			expectedPos: &point.Point{X: 1, Y: 2},
			expectedErr: nil,
		},
		{
			name: "should ignore move command that would push robot out of the board",
			tbl: func() *table.Table {
				tbl := table.New(5, 5)
				tbl.PlaceRobot(point.Point{X: 0, Y: 1}, direction.West)
				return tbl
			},
			expectedPos: &point.Point{X: 0, Y: 1},
			expectedErr: table.ErrEndingPositionOutOfBounds,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual, err := tt.tbl().MoveRobot()
			require.Equal(t, tt.expectedPos, actual)
			require.Equal(t, tt.expectedErr, err)
		})
	}
}

func TestRotateRobot(t *testing.T) {
	t.Parallel()

	tests := [...]struct {
		name           string
		tbl            func() *table.Table
		rotateLeft     bool
		expectedFacing *direction.Direction
		expectedErr    error
	}{
		{
			name: "should ignore rotate command when placement was not done",
			tbl: func() *table.Table {
				tbl := table.New(5, 5)
				return tbl
			},
			rotateLeft:     true,
			expectedFacing: nil,
			expectedErr:    table.ErrUninitializedPlacement,
		},
		{
			name: "should successfully rotate robot left",
			tbl: func() *table.Table {
				tbl := table.New(5, 5)
				tbl.PlaceRobot(point.Point{X: 1, Y: 1}, direction.North)
				return tbl
			},
			rotateLeft:     true,
			expectedFacing: &direction.West,
			expectedErr:    nil,
		},
		{
			name: "should successfully rotate robot right",
			tbl: func() *table.Table {
				tbl := table.New(5, 5)
				tbl.PlaceRobot(point.Point{X: 1, Y: 1}, direction.South)
				return tbl
			},
			rotateLeft:     false,
			expectedFacing: &direction.West,
			expectedErr:    nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual, err := tt.tbl().RotateRobot(tt.rotateLeft)
			require.Equal(t, tt.expectedFacing, actual)
			require.Equal(t, tt.expectedErr, err)
		})
	}
}

func TestReport(t *testing.T) {

	reportBuf := bytes.NewBufferString("")

	tests := [...]struct {
		name        string
		tbl         func() *table.Table
		expected    string
		expectedErr error
	}{
		{
			name: "should ignore report command when placement was not done",
			tbl: func() *table.Table {
				tbl := table.New(5, 5, table.WithReportOutput(reportBuf))
				return tbl
			},
			expected:    "",
			expectedErr: table.ErrUninitializedPlacement,
		},
		{
			name: "should successfully report robots position and facing",
			tbl: func() *table.Table {
				tbl := table.New(5, 5, table.WithReportOutput(reportBuf))
				tbl.PlaceRobot(point.Point{X: 2, Y: 3}, direction.East)
				return tbl
			},
			expected:    "Robot position: (2, 3) facing: EAST\n",
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			reportBuf.Reset()
			err := tt.tbl().Report()
			require.Equal(t, tt.expected, reportBuf.String())
			require.Equal(t, tt.expectedErr, err)
		})
	}
}
