package main

import (
	"sync"
)

type ElevatorController struct {
	Elevators []*Elevator
}

func NewElevatorController(numElevators int, wg *sync.WaitGroup) *ElevatorController {
	elevators := make([]*Elevator, numElevators)
	for i := range elevators {
		elevators[i] = &Elevator{
			ID:        i + 1,
			RequestCh: make(chan Request),
			Wg:        wg,
		}
		wg.Add(1)
		go elevators[i].Run()
	}
	return &ElevatorController{Elevators: elevators}
}

func (ec *ElevatorController) findNearestElevator(source int) int {
	minDist := 1 << 30
	chosen := 0
	for i, e := range ec.Elevators {
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

func (ec *ElevatorController) DispatchRequests(requests []Request) {
	for _, req := range requests {
		chosen := ec.findNearestElevator(req.Source)
		ec.Elevators[chosen].RequestCh <- req
	}
	for _, e := range ec.Elevators {
		close(e.RequestCh)
	}
}
