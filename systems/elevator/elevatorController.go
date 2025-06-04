package main

import "fmt"

type ElevatorController struct {
	Elevators []Elevator
}

func NewElevatorController(numElevators int) *ElevatorController {
	elevators := make([]Elevator, numElevators)
	for i := range elevators {
		elevators[i] = Elevator{ID: i + 1, CurrentFloor: 0, Direction: Up}
	}
	return &ElevatorController{Elevators: elevators}
}

func (ec *ElevatorController) AssignRequest(req Request) int {
	minDist := 1 << 30
	chosen := 0
	for i, e := range ec.Elevators {
		dist := e.CurrentFloor - req.Source
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

func (ec *ElevatorController) HandleRequests(requests []Request) {
	for i, req := range requests {
		chosen := ec.AssignRequest(req)
		e := &ec.Elevators[chosen]
		fmt.Printf("\nProcessing request %d with Elevator %d:\n", i+1, e.ID)
		if e.CurrentFloor != req.Source {
			fmt.Printf("Elevator %d moving to pick up at floor %d\n", e.ID, req.Source)
			e.MoveToFloor(req.Source)
		}
		e.Direction = req.Direction
		e.MoveToFloor(req.Destination)
	}
}
