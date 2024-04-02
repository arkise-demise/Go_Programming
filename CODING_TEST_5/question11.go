package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
    ID       int
    Duration time.Duration
}

func executeTask(task Task, mutexes map[int]*sync.Mutex, done chan<- int) {
    
    mutexes[task.ID].Lock() 
    defer mutexes[task.ID].Unlock() 
    fmt.Printf("Task %d started\n", task.ID)
    time.Sleep(task.Duration) 
    fmt.Printf("Task %d finished\n", task.ID)

    done <- task.ID 
}
func scheduleTasks(tasks []Task) {
    mutexes := make(map[int]*sync.Mutex)
    done := make(chan int)

    for _, task := range tasks {
        if _, exists := mutexes[task.ID]; !exists {
            mutexes[task.ID] = &sync.Mutex{}
        }
    }

    for _, task := range tasks {
        go executeTask(task, mutexes, done)
    }


    for range tasks {
        <-done
    }
}

func main() {
    tasks := []Task{
        {ID: 1, Duration: 2 * time.Second},
        {ID: 2, Duration: 1 * time.Second},
        {ID: 3, Duration: 3 * time.Second},
    }

    scheduleTasks(tasks)
}
