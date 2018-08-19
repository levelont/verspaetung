package coordinate

import (
	"strconv"
)

type Coordinate struct {
	x int
	y int
}

func New(x string, y string) (coordinate Coordinate, err error) {
	var xValue, yValue int
	xValue, err = strconv.Atoi(x)
	if err != nil {
		return
	}
	yValue, err = strconv.Atoi(y)
	if err != nil {
		return
	}
	coordinate = Coordinate{x: xValue, y: yValue}
	err = nil
	return
}
