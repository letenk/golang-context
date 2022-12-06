package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	deadline := time.Now().Add(1 * time.Second)
	ctxCancel, cancel := context.WithDeadline(ctx, deadline)
	defer cancel()

	// Create new value
	greetContext := context.WithValue(ctxCancel, "myKey", "Hola")
	// Call function greetings and sent greetContext
	go greetings(greetContext)

	// Get value from context
	val := ctx.Value("myKey")
	// Print value
	fmt.Println("Value in context says:", val)
	fmt.Printf("doSomething: finished \n")
	time.Sleep(2 * time.Second)
}

func greetings(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Printf("greetings err: %s\n", err)
			}
			fmt.Printf("greetings: finished \n")
			return
		default:
			// Get value from context
			val := ctx.Value("myKey")
			// Print value
			fmt.Println("Value context in function greetings says:", val)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	// Create new context
	ctx := context.Background()
	// Add value to context
	ctx = context.WithValue(ctx, "myKey", "Hello")
	// Call function doSomething and sent context
	doSomething(ctx)
}
