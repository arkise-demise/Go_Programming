package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a context with a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Channel to signal task completion
	done := make(chan struct{})

	// Launch a goroutine to perform the task
	go func() {
		// Simulate task taking 3 seconds to complete
		time.Sleep(3 * time.Second)
		// Signal task completion
		done <- struct{}{}
	}()

	select {
	case <-done:
		fmt.Println("Task completed")
	case <-ctx.Done():
		fmt.Println("Context canceled")
	}
}
