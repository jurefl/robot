package command_test

import (
	"robot/internal/direction"
	"robot/internal/point"
)

type tableMock struct {
	placeRobotFn  func(pos point.Point, facing direction.Direction) error
	rotateRobotFn func(left bool) (*direction.Direction, error)
	moveRobotFn   func() (*point.Point, error)
	reportFn      func() error
	fnCnt         map[string]int
}

func (m *tableMock) funcCallCountInc(funcName string) {
	if m.fnCnt == nil {
		m.fnCnt = map[string]int{}
	}
	m.fnCnt[funcName]++
}

func (m *tableMock) PlaceRobot(pos point.Point, facing direction.Direction) error {
	m.funcCallCountInc("PlaceRobot")
	if m.placeRobotFn != nil {
		return m.placeRobotFn(pos, facing)
	}
	return nil
}

func (m *tableMock) RotateRobot(left bool) (*direction.Direction, error) {
	m.funcCallCountInc("RotateRobot")
	if m.rotateRobotFn != nil {
		return m.rotateRobotFn(left)
	}
	return &direction.North, nil
}

func (m *tableMock) MoveRobot() (*point.Point, error) {
	m.funcCallCountInc("MoveRobot")
	if m.moveRobotFn != nil {
		return m.moveRobotFn()
	}
	return &point.Point{X: 0, Y: 0}, nil
}

func (m *tableMock) Report() error {
	m.funcCallCountInc("Report")
	if m.reportFn != nil {
		return m.reportFn()
	}
	return nil
}
