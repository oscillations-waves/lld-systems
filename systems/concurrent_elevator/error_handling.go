package main

import "fmt"

func ValidateElevatorCount(numElevators int, err error) bool {
	if err != nil || numElevators <= 0 {
		fmt.Println("Error: Please enter a valid positive integer for number of elevators.")
		return false
	}
	return true
}

func ValidateRequestCount(n int, err error) bool {
	if err != nil || n <= 0 {
		fmt.Println("Error: Please enter a valid positive integer for number of requests.")
		return false
	}
	return true
}

func ValidateFloor(floor, minFloor, maxFloor int, label string) bool {
	if floor < minFloor || floor > maxFloor {
		fmt.Printf("Error: %s floor must be between %d and %d.\n", label, minFloor, maxFloor)
		return false
	}
	return true
}

func ValidateDirection(dir string, err error) (Direction, bool) {
	if err != nil || (dir != "up" && dir != "down") {
		fmt.Println("Error: Direction must be 'up' or 'down'.")
		return 0, false
	}
	if dir == "up" {
		return Up, true
	}
	return Down, true
}

func ValidateSourceDest(source, dest int) bool {
	if source == dest {
		fmt.Println("Error: Source and destination floors cannot be the same.")
		return false
	}
	return true
}
