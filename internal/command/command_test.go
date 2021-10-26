package command_test

import (
	"robot/internal/command"
	"robot/internal/direction"
	"robot/internal/point"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnmarshal(t *testing.T) {
	t.Parallel()

	tests := [...]struct {
		name          string
		tbl           *tableMock
		command       string
		expectedFnCnt map[string]int
		shouldErr     bool
	}{
		{
			name: "should unmarshal place command",
			tbl: &tableMock{
				placeRobotFn: func(pos point.Point, facing direction.Direction) error {
					require.Equal(t, point.Point{X: 1, Y: 2}, pos)
					require.Equal(t, direction.North, facing)
					return nil
				},
			},
			command:       "PLACE 1,2,NORTH",
			expectedFnCnt: map[string]int{"PlaceRobot": 1},
			shouldErr:     false,
		},
		{
			name: "should fail to unmarshal invalid place command",
			tbl: &tableMock{
				placeRobotFn: func(pos point.Point, facing direction.Direction) error {
					require.Equal(t, point.Point{X: 1, Y: 2}, pos)
					require.Equal(t, direction.North, facing)
					return nil
				},
			},
			command:       "PLACE one,2,NORTH",
			expectedFnCnt: map[string]int{},
			shouldErr:     true,
		},
		{
			name:          "should unmarshal move command",
			tbl:           &tableMock{},
			command:       "MOVE",
			expectedFnCnt: map[string]int{"MoveRobot": 1},
			shouldErr:     false,
		},
		{
			name: "should unmarshal left command",
			tbl: &tableMock{
				rotateRobotFn: func(left bool) (*direction.Direction, error) {
					require.Equal(t, true, left)
					return nil, nil
				},
			},
			command:       "LEFT",
			expectedFnCnt: map[string]int{"RotateRobot": 1},
			shouldErr:     false,
		},
		{
			name: "should unmarshal right command",
			tbl: &tableMock{
				rotateRobotFn: func(left bool) (*direction.Direction, error) {
					require.Equal(t, false, left)
					return nil, nil
				},
			},
			command:       "RIGHT",
			expectedFnCnt: map[string]int{"RotateRobot": 1},
			shouldErr:     false,
		},
		{
			name:          "should unmarshal report command",
			tbl:           &tableMock{},
			command:       "REPORT",
			expectedFnCnt: map[string]int{"Report": 1},
			shouldErr:     false,
		},
		{
			name: "should fail to unmarshal unknown command",
			tbl: &tableMock{
				rotateRobotFn: func(left bool) (*direction.Direction, error) {
					require.Equal(t, false, left)
					return nil, nil
				},
			},
			command:       "UNKNOWN",
			expectedFnCnt: map[string]int{},
			shouldErr:     true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var cmd command.Command
			err := cmd.Unmarshal(tt.command)
			if tt.shouldErr {
				require.Error(t, err)
				return
			}

			cmd(tt.tbl)
			require.Equal(t, tt.expectedFnCnt, tt.tbl.fnCnt)
		})
	}
}
