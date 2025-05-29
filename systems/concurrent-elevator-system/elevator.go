package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
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
	if !ValidateElevatorCount(numElevators, err) {
		return
	}

	var wg sync.WaitGroup
	controller := NewElevatorController(numElevators, &wg)

	// Graceful shutdown setup
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-shutdown
		fmt.Println("\nGraceful shutdown initiated. Closing elevator channels...")
		for _, e := range controller.Elevators {
			close(e.RequestCh)
		}
		os.Exit(0)
	}()

	var n int
	fmt.Print("Enter number of requests: ")
	_, err = fmt.Scan(&n)
	if !ValidateRequestCount(n, err) {
		return
	}
	requests := make([]Request, 0, n)
	requestSet := make(map[string]struct{})
	var minFloor, maxFloor int = -5, 100 // Set building floor range here
	for i := 0; i < n; i++ {
		var source, dest int
		var dir string
		fmt.Printf("Request %d - Enter source floor: ", i+1)
		_, err = fmt.Scan(&source)
		if err != nil {
			fmt.Println("Error: Invalid input for source floor.")
			return
		}
		if !ValidateFloor(source, minFloor, maxFloor, "Source") {
			return
		}
		fmt.Printf("Request %d - Enter destination floor: ", i+1)
		_, err = fmt.Scan(&dest)
		if err != nil {
			fmt.Println("Error: Invalid input for destination floor.")
			return
		}
		if !ValidateFloor(dest, minFloor, maxFloor, "Destination") {
			return
		}
		fmt.Printf("Request %d - Enter direction (up/down): ", i+1)
		_, err = fmt.Scan(&dir)
		direction, valid := ValidateDirection(dir, err)
		if !valid {
			return
		}
		if !ValidateSourceDest(source, dest) {
			return
		}
		key := fmt.Sprintf("%d-%d-%d", source, dest, direction)
		if _, exists := requestSet[key]; exists {
			fmt.Printf("Warning: Duplicate request from %d to %d (%s) ignored.\n", source, dest, dir)
			continue
		}
		requestSet[key] = struct{}{}
		requests = append(requests, Request{Source: source, Destination: dest, Direction: direction})
	}

	if len(requests) == 0 {
		fmt.Println("No valid requests to process. Exiting.")
		return
	}

	// Starvation/Deadlock detection: If all elevators are busy and requests remain unprocessed for too long, warn the user.
	// We'll use a timer to simulate a simple starvation detection for demonstration.
	starvationTimer := time.AfterFunc(10*time.Second, func() {
		fmt.Println("Warning: Some requests are taking too long to process. Possible deadlock or starvation detected.")
	})

	controller.DispatchRequests(requests)
	wg.Wait()
	starvationTimer.Stop()
}
