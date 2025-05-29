package main

import (
	"fmt"
)

type Direction int

const (
	Up Direction = iota
	Down
)

type Elevator struct {
	ID           int
	CurrentFloor int
	Direction    Direction
}

type Request struct {
	Source      int
	Destination int
	Direction   Direction
}

func (e *Elevator) MoveToFloor(dest int) {
	if dest == e.CurrentFloor {
		fmt.Printf("Elevator %d is already at floor %d\n", e.ID, dest)
		return
	}
	if dest > e.CurrentFloor {
		e.Direction = Up
		fmt.Printf("Elevator %d going up from %d to %d\n", e.ID, e.CurrentFloor, dest)
	} else {
		e.Direction = Down
		fmt.Printf("Elevator %d going down from %d to %d\n", e.ID, e.CurrentFloor, dest)
	}
	for e.CurrentFloor != dest {
		if e.Direction == Up {
			e.CurrentFloor++
		} else {
			e.CurrentFloor--
		}
		fmt.Printf("Elevator %d passing floor %d\n", e.ID, e.CurrentFloor)
	}
	fmt.Printf("Elevator %d arrived at floor %d\n", e.ID, e.CurrentFloor)
}

func findNearestElevator(elevators []Elevator, source int) int {
	minDist := 1 << 30
	chosen := 0
	for i, e := range elevators {
		dist := e.CurrentFloor - source
		if dist < 0 {
			dist = -dist
		}
		if dist < minDist {
			minDist = dist
			chosen = i
		}
	}
	return chosen
}

func ProcessRequests(elevators []Elevator, requests []Request) {
	for i, req := range requests {
		chosen := findNearestElevator(elevators, req.Source)
		e := &elevators[chosen]
		fmt.Printf("\nProcessing request %d with Elevator %d:\n", i+1, e.ID)
		if e.CurrentFloor != req.Source {
			fmt.Printf("Elevator %d moving to pick up at floor %d\n", e.ID, req.Source)
			e.MoveToFloor(req.Source)
		}
		e.Direction = req.Direction
		e.MoveToFloor(req.Destination)
	}
}

func main() {
	var numElevators int
	fmt.Print("Enter number of elevators: ")
	fmt.Scan(&numElevators)
	elevators := make([]Elevator, numElevators)
	for i := range elevators {
		elevators[i] = Elevator{ID: i + 1, CurrentFloor: 0, Direction: Up}
	}

	var n int
	fmt.Print("Enter number of requests: ")
	fmt.Scan(&n)
	requests := make([]Request, n)
	for i := 0; i < n; i++ {
		var source, dest int
		var dir string
		fmt.Printf("Request %d - Enter source floor: ", i+1)
		fmt.Scan(&source)
		fmt.Printf("Request %d - Enter destination floor: ", i+1)
		fmt.Scan(&dest)
		fmt.Printf("Request %d - Enter direction (up/down): ", i+1)
		fmt.Scan(&dir)
		var direction Direction
		if dir == "up" {
			direction = Up
		} else {
			direction = Down
		}
		requests[i] = Request{Source: source, Destination: dest, Direction: direction}
	}

	ProcessRequests(elevators, requests)
}
