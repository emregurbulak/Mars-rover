package validations

import (
	"plateau-rovers/constants"
	"strconv"
	"strings"
)

func IsPlateauValid(val string) bool {
	trimmedVal := strings.ReplaceAll(val, " ", "")
	trimmedVal = strings.TrimSpace(trimmedVal)
	_, err := strconv.Atoi(trimmedVal)
	if err != nil {
		return false
	}

	plateauSizeArr := strings.Fields(val)
	if len(plateauSizeArr) != 2 {
		return false
	}

	xSize, _ := strconv.Atoi(plateauSizeArr[0])
	ySize, _ := strconv.Atoi(plateauSizeArr[1])

	if xSize == 0 || ySize == 0 {
		return false
	}

	return true
}

func IsRoverCoordinates(val string, maxXCoordinate, maxYCoordinate int) (resp bool, errMsg string) {
	roverCoordinates := strings.Fields(val)

	// like [ 5 5 N ]
	if len(roverCoordinates) != 3 {
		return false, constants.RoverCoordinatsDoesNotCorrectFormat
	}

	readDirection := roverCoordinates[2] // Third input shoud be a direction
	isDirectionValid := false
	if readDirection == constants.East || readDirection == constants.West ||
		readDirection == constants.South || readDirection == constants.North {
		isDirectionValid = true
	}

	if !isDirectionValid {
		return false, constants.RoverDirectionIsNotValid
	}

	xCoordinate, _ := strconv.Atoi(roverCoordinates[0])
	yCoordinate, _ := strconv.Atoi(roverCoordinates[1])

	if xCoordinate > maxXCoordinate || yCoordinate > maxYCoordinate {
		return false, constants.RoverCoordinatesExceededPlateau
	}

	return true, ""
}

func IsRoverDirectionValid(val string) (resp bool, errMsg string) {
	isRoverDirectionValid := true
	trimmedVal := strings.ReplaceAll(val, " ", "")
	roverDirections := strings.TrimSpace(trimmedVal)
	if len(roverDirections) == 0 {
		return false, constants.RoverMovementDirectionInfo
	}

	for key := 0; key < len(roverDirections); key++ {
		directionKey := string(roverDirections[key])
		if directionKey != "L" && directionKey != "R" && directionKey != "M" {
			isRoverDirectionValid = false
			break
		}
	}

	if !isRoverDirectionValid {
		return false, constants.RoverMovementDirectionIsNotValid
	}
	return true, ""

}
