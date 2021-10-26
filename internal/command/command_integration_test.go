package command_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	"robot/internal/command"
	"robot/internal/table"
)

func TestScanCommandList(t *testing.T) {
	t.Parallel()

	tests := [...]struct {
		name          string
		tbl           *tableMock
		commandFile   string
		expectedFnCnt map[string]int
		shouldErr     bool
	}{
		{
			name:        "should successfully scan commands to draw letter y",
			tbl:         &tableMock{},
			commandFile: "./fixtures/y.txt",
			expectedFnCnt: map[string]int{
				"PlaceRobot":  1,
				"MoveRobot":   10,
				"RotateRobot": 6,
				"Report":      1,
			},
			shouldErr: false,
		},
		{
			name:        "should successfully scan commands to draw letter u",
			tbl:         &tableMock{},
			commandFile: "./fixtures/u.txt",
			expectedFnCnt: map[string]int{
				"PlaceRobot":  1,
				"MoveRobot":   10,
				"RotateRobot": 2,
				"Report":      1,
			},
			shouldErr: false,
		},
		{
			name:        "should successfully scan commands to draw letter m",
			tbl:         &tableMock{},
			commandFile: "./fixtures/m.txt",
			expectedFnCnt: map[string]int{
				"PlaceRobot":  1,
				"MoveRobot":   15,
				"RotateRobot": 10,
				"Report":      1,
			},
			shouldErr: false,
		},
		{
			name:          "should fail to scan invalid commands",
			tbl:           &tableMock{},
			commandFile:   "./fixtures/invalid.txt",
			expectedFnCnt: map[string]int{},
			shouldErr:     true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			cmds, err := command.ScanCommandList(tt.commandFile)
			if tt.shouldErr {
				require.Error(t, err)
				return
			}

			for _, cmd := range cmds {
				cmd(tt.tbl)
			}
			require.EqualValues(t, tt.expectedFnCnt, tt.tbl.fnCnt)
		})
	}
}

func TestIntegration(t *testing.T) {
	tests := [...]struct {
		name           string
		commandFile    string
		expectedReport string
	}{
		{
			name:           "should successfully scan commands to draw letter y",
			commandFile:    "./fixtures/y.txt",
			expectedReport: "Robot position: (2, 0) facing: SOUTH\n",
		},
		{
			name:           "should successfully scan commands to draw letter u",
			commandFile:    "./fixtures/u.txt",
			expectedReport: "Robot position: (2, 4) facing: NORTH\n",
		},
		{
			name:           "should successfully scan commands to draw letter m",
			commandFile:    "./fixtures/m.txt",
			expectedReport: "Robot position: (3, 0) facing: SOUTH\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reportBuf := bytes.NewBufferString("")
			tbl := table.New(5, 5, table.WithReportOutput(reportBuf))
			cmds, _ := command.ScanCommandList(tt.commandFile)
			for _, cmd := range cmds {
				cmd(tbl)
			}
			require.EqualValues(t, tt.expectedReport, reportBuf.String())
		})
	}
}
