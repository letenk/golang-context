package main

import (
	"context"
	"fmt"
)

func doSomething(ctx context.Context) {
	fmt.Println("Hello from context.")
}

func main() {
	ctx := context.Background()
	doSomething(ctx)
}
