package directions

import (
	"plateau-rovers/constants"
)

//input (N,W,S,E), (L R) , output: (N,W,S,E)
func FindRoversNewDirection(direction, rotate string) (newDirection string) {
	var newDirectionNum int
	directionNum, _ := constants.DirectionRotateNumbers[direction]

	if rotate == "L" && directionNum != constants.DirectionRotateNumbers["E"] {
		newDirectionNum = directionNum + 1
	} else if rotate == "L" && directionNum == constants.DirectionRotateNumbers["E"] {
		newDirectionNum = constants.DirectionRotateNumbers["N"]
	} else if directionNum == constants.DirectionRotateNumbers["N"] {
		newDirectionNum = constants.DirectionRotateNumbers["E"]
	} else {
		newDirectionNum = directionNum - 1
	}
	for key, value := range constants.DirectionRotateNumbers {
		if value == newDirectionNum {
			newDirection = key
		}
	}
	return newDirection
}
