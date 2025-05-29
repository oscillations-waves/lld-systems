package main

import (
	"fmt"
	"sync"
)

type Direction int

const (
	Up Direction = iota
	Down
)

type Request struct {
	Source      int
	Destination int
	Direction   Direction
}

type Elevator struct {
	ID           int
	CurrentFloor int
	Direction    Direction
	RequestCh    chan Request
	Wg           *sync.WaitGroup
}

func (e *Elevator) Run() {
	defer e.Wg.Done()
	for req := range e.RequestCh {
		fmt.Printf("\nElevator %d received request: from %d to %d (%v)\n", e.ID, req.Source, req.Destination, req.Direction)
		if e.CurrentFloor != req.Source {
			fmt.Printf("Elevator %d moving to pick up at floor %d\n", e.ID, req.Source)
			e.moveToFloor(req.Source)
		}
		e.Direction = req.Direction
		e.moveToFloor(req.Destination)
	}
}

func (e *Elevator) moveToFloor(dest int) {
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

func main() {
	var numElevators int
	fmt.Print("Enter number of elevators: ")
	_, err := fmt.Scan(&numElevators)
	if err != nil || numElevators <= 0 {
		fmt.Println("Error: Please enter a valid positive integer for number of elevators.")
		return
	}

	var wg sync.WaitGroup
	controller := NewElevatorController(numElevators, &wg)

	var n int
	fmt.Print("Enter number of requests: ")
	_, err = fmt.Scan(&n)
	if err != nil || n <= 0 {
		fmt.Println("Error: Please enter a valid positive integer for number of requests.")
		return
	}
	requests := make([]Request, n)
	for i := 0; i < n; i++ {
		var source, dest int
		var dir string
		fmt.Printf("Request %d - Enter source floor: ", i+1)
		_, err = fmt.Scan(&source)
		if err != nil {
			fmt.Println("Error: Invalid input for source floor.")
			return
		}
		fmt.Printf("Request %d - Enter destination floor: ", i+1)
		_, err = fmt.Scan(&dest)
		if err != nil {
			fmt.Println("Error: Invalid input for destination floor.")
			return
		}
		fmt.Printf("Request %d - Enter direction (up/down): ", i+1)
		_, err = fmt.Scan(&dir)
		if err != nil || (dir != "up" && dir != "down") {
			fmt.Println("Error: Direction must be 'up' or 'down'.")
			return
		}
		if source == dest {
			fmt.Println("Error: Source and destination floors cannot be the same.")
			return
		}
		var direction Direction
		if dir == "up" {
			direction = Up
		} else {
			direction = Down
		}
		requests[i] = Request{Source: source, Destination: dest, Direction: direction}
	}

	controller.DispatchRequests(requests)
	wg.Wait()
}
