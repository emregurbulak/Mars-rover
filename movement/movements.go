package movement

import "plateau-rovers/constants"

func FindNewCoordinates(xCoordinate, yCoordinate int, direction string) (newXcoordinate, newYcoordinate int) {
	newXcoordinate = xCoordinate
	newYcoordinate = yCoordinate
	switch direction {
	case constants.North:
		newYcoordinate = yCoordinate + 1
	case constants.West:
		newXcoordinate = xCoordinate - 1
	case constants.East:
		newXcoordinate = xCoordinate + 1
	case constants.South:
		newYcoordinate = yCoordinate - 1
	}

	return newXcoordinate, newYcoordinate
}
