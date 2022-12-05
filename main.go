package main

import (
	"context"
	"fmt"
)

func doSomething(ctx context.Context) {
	// Get value from context
	val := ctx.Value("myKey")
	// Print value
	fmt.Println("Value in context says:", val)
}

func main() {
	// Create new context
	ctx := context.Background()
	// Add value to context
	ctx = context.WithValue(ctx, "myKey", "Hello")
	// Call function doSomething and sent context
	doSomething(ctx)
}
