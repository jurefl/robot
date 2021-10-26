package command

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"robot/internal/direction"
	"robot/internal/point"
)

type Table interface {
	PlaceRobot(pos point.Point, facing direction.Direction) error
	RotateRobot(left bool) (*direction.Direction, error)
	MoveRobot() (*point.Point, error)
	Report() error
}

// Command that can be executed against robot table
type Command func(t Table)

var (
	leftCmd Command = func(t Table) {
		t.RotateRobot(true)
	}

	rightCmd Command = func(t Table) {
		t.RotateRobot(false)
	}

	moveCmd Command = func(t Table) {
		t.MoveRobot()
	}

	reportCmd Command = func(t Table) {
		t.Report()
	}
)

// ScanCommandList parses commands from the text file
func ScanCommandList(fileName string) ([]Command, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	cmdList := []Command{}
	for scanner.Scan() {
		var cmd Command
		if err := cmd.Unmarshal(scanner.Text()); err != nil {
			return nil, err
		}
		cmdList = append(cmdList, cmd)
	}

	return cmdList, nil
}

// Unmarshal deserialize individual command
func (c *Command) Unmarshal(s interface{}) error {
	src, ok := s.(string)
	if !ok {
		return errors.New("non string source data types not supported")
	}

	txtCmd := strings.Fields(strings.TrimSpace(strings.ToUpper(src)))
	if len(txtCmd) == 0 {
		return fmt.Errorf("empty command detected: '%s'", src)
	}

	switch txtCmd[0] {
	case "PLACE":
		var cmdParams []string
		if len(txtCmd) > 1 {
			cmdParams = strings.Split(txtCmd[1], ",")
		}
		if len(cmdParams) < 3 {
			return fmt.Errorf("PLACE command requires 3 parameters, but %d were detected: '%s'", len(cmdParams), src)
		}
		cmd, err := placeCmd(cmdParams[0], cmdParams[1], cmdParams[2])
		if err != nil {
			return err
		}
		*c = cmd

	case "LEFT":
		*c = leftCmd
	case "RIGHT":
		*c = rightCmd
	case "MOVE":
		*c = moveCmd
	case "REPORT":
		*c = reportCmd
	default:
		return fmt.Errorf("invalid command detected: '%s'", src)
	}
	return nil
}

// placecmd deserialize place command
func placeCmd(x, y, drctn string) (Command, error) {
	posX, err := strconv.Atoi(x)
	if err != nil {
		return nil, fmt.Errorf("x pos parameters not a number(%s): %s", x, err.Error())
	}
	posY, err := strconv.Atoi(y)
	if err != nil {
		return nil, fmt.Errorf("y pos parameters not a number(%s): %s", y, err.Error())
	}

	var d direction.Direction
	switch drctn {
	case "EAST":
		d = direction.East
	case "NORTH":
		d = direction.North
	case "WEST":
		d = direction.West
	case "SOUTH":
		d = direction.South
	default:
		return nil, fmt.Errorf("invalid direction parameter detected: '%s'", drctn)
	}

	return func(t Table) {
		t.PlaceRobot(point.Point{X: posX, Y: posY}, d)
	}, nil
}
