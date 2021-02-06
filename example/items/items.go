package main

import (
	"context"

	mc "github.com/brittonhayes/go-minecraft"
	"github.com/brittonhayes/go-minecraft/items"
)

func main() {
	c := mc.NewClient("localhost:1234", "abc1234")

	var ctx context.Context
	_, _ = c.Give(ctx, "", items.AcaciaBoat, 5)
}
