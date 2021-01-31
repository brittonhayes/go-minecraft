package main

import (
	"context"
	"fmt"
	minecraft "github.com/brittonhayes/go-minecraft"
)

func main() {
	// Initialize client
	c := minecraft.NewClient("localhost:1234", "password")

	// Setup context for request
	ctx := context.Background()

	// Give the player items
	res, err := c.Give(ctx, "johndoe", minecraft.Bedrock, 5)
	if err != nil {
		panic(err)
	}

	// Print out the response
	fmt.Println(res)
}
