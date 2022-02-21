package cartesian

import (
	"github.com/leosimoesp/civi-backend-exercise/internal/app/datatype"
)

func ManhattanDistance(p1, p2 datatype.Point) int {
	return Abs(p1.X-p2.X) + Abs(p1.Y-p2.Y)
}

func Abs(value int) int {
	if value < 0 {
		return value * -1
	}
	return value
}

func CheckPointInManhattanDistance(origin, p2 datatype.Point, targetDistance int) (bool, int) {

	distance := ManhattanDistance(origin, p2)

	//if distance is equal to zero then origin and target are the same points
	return distance <= targetDistance, distance
}
