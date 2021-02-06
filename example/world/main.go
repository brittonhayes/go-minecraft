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

	// Set the world weather to rain
	res := c.Weather(ctx, minecraft.WeatherRain)

	// Print out the response
	fmt.Println(res)
}
