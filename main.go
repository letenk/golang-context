package main

import (
	"context"
	"fmt"
)

func doSomething(ctx context.Context) {
	// Create new value
	greetContext := context.WithValue(ctx, "myKey", "Hola")
	// Call function greetings and sent greetContext
	greetings(greetContext)

	// Get value from context
	val := ctx.Value("myKey")
	// Print value
	fmt.Println("Value in context says:", val)
}

func greetings(ctx context.Context) {
	// Get value from context
	val := ctx.Value("myKey")
	// Print value
	fmt.Println("Value context in function greetings says:", val)
}

func main() {
	// Create new context
	ctx := context.Background()
	// Add value to context
	ctx = context.WithValue(ctx, "myKey", "Hello")
	// Call function doSomething and sent context
	doSomething(ctx)
}
