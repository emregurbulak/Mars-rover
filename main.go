package main

import (
	"bufio"
	"fmt"
	"os"
	"plateau-rovers/constants"
	"plateau-rovers/directions"
	"plateau-rovers/movement"
	"plateau-rovers/validations"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	isPlateauValid := false

	errMsg := ""
	var plateauSizeTxt string
	fmt.Print("Please enter the plateau size:")
	for !isPlateauValid {
		plateauSizeTxt, _ = reader.ReadString('\n')
		isPlateauValid = validations.IsPlateauValid(plateauSizeTxt)
		if !isPlateauValid {
			fmt.Println(constants.PlateauSizeDoesNotCorrectFormat)
		}
	}
	maxXCoordinate, _ := strconv.Atoi(strings.Fields(plateauSizeTxt)[0])
	maxYCoordinate, _ := strconv.Atoi(strings.Fields(plateauSizeTxt)[1])
	isPlateauDiscoveryEnough := false
	for !isPlateauDiscoveryEnough {
		isRoverCoordinatesValid := false
		fmt.Print("Please enter the rover coordinates:")
		var roverCoordinatesText string
		for !isRoverCoordinatesValid {
			roverCoordinatesText, _ = reader.ReadString('\n')
			isRoverCoordinatesValid, errMsg = validations.IsRoverCoordinates(roverCoordinatesText, maxXCoordinate, maxYCoordinate)
			if !isRoverCoordinatesValid {
				fmt.Println(errMsg)
			}
		}

		xCoordinate, _ := strconv.Atoi(strings.Fields(roverCoordinatesText)[0])
		yCoordinate, _ := strconv.Atoi(strings.Fields(roverCoordinatesText)[1])
		direction := strings.Fields(roverCoordinatesText)[2]

		addCoordinate := true
		for addCoordinate {
			isRoverDirectionValid := false
			fmt.Print("Please enter the direction:")
			var directionText string
			for !isRoverDirectionValid {
				directionText, _ = reader.ReadString('\n')
				isRoverDirectionValid, errMsg = validations.IsRoverDirectionValid(directionText)
				if !isRoverDirectionValid {
					fmt.Println(errMsg)
				}
			}

			trimmedVal := strings.ReplaceAll(directionText, " ", "")
			roverRotateAndMovingInfo := strings.TrimSpace(trimmedVal)

			newDirection := direction
			isNewCoordinateValid := true
			for i := 0; i < len([]rune(roverRotateAndMovingInfo)); i++ {
				currentVal := string(roverRotateAndMovingInfo[i])
				if currentVal == constants.Left || currentVal == constants.Right {
					newDirection = directions.FindRoversNewDirection(direction, string(roverRotateAndMovingInfo[i]))
					direction = newDirection

				} else {
					newXCoordinate, newYCoordinate := movement.FindNewCoordinates(xCoordinate, yCoordinate, newDirection)
					if newXCoordinate > maxXCoordinate || newYCoordinate > maxYCoordinate ||
						newXCoordinate < 0 || newYCoordinate < 0 {
						fmt.Println()
						isNewCoordinateValid = false
						break

					} else {
						xCoordinate = newXCoordinate
						yCoordinate = newYCoordinate

					}
				}

			}
			if !isNewCoordinateValid {
				addCoordinate = true
				fmt.Print("New coordinates are exeeded plateau.")
			} else {
				addCoordinate = false
				fmt.Printf("%d %d %s: ", xCoordinate, yCoordinate, direction)
			}
		}

		fmt.Print("\n Do you want to discover with another rover: ( Y / N) :")
		isDiscoverChoosingValid := false
		for !isDiscoverChoosingValid {
			chooseAnotherRover, _ := reader.ReadString('\n')
			val := strings.ReplaceAll(chooseAnotherRover, " ", "")
			trimmedVal := strings.TrimSpace(val)
			if trimmedVal != "Y" && trimmedVal != "N" {
				fmt.Print("\n Please choose one: ( Y / N) :")
				isDiscoverChoosingValid = false
			} else {
				if trimmedVal == "Y" {
					isPlateauDiscoveryEnough = false
				} else {
					isPlateauDiscoveryEnough = true
				}
				isDiscoverChoosingValid = true
			}
		}
	}
	fmt.Print("\n Thank you for this journey.All is well")
}
