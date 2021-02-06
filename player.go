package go_minecraft

import (
	"context"
	"fmt"

	"github.com/brittonhayes/go-minecraft/internal/rcon"
	"github.com/brittonhayes/go-minecraft/items"
)

var _ Player = (*PlayerService)(nil)

// Player is the primary interface for the all player-related commands
type Player interface {
	Give(ctx context.Context, name string, item items.Item, quantity int) (string, error)
	List(ctx context.Context) string
	Meetup(ctx context.Context, player1, player2 string) string
	Teleport(ctx context.Context, player, x, y, z string) string
	Kill(ctx context.Context, player string) string
}

type PlayerService struct {
	addr     string
	password string
}

func (p *PlayerService) List(ctx context.Context) string {
	c := rcon.NewConnection(ctx, p.addr, p.password)
	defer c.Close()

	result, err := c.SendCommand("list")
	if err != nil {
		return err.Error()
	}

	return result
}

// Kill kills the player with the provided name
func (p *PlayerService) Kill(ctx context.Context, player string) string {
	c := rcon.NewConnection(ctx, p.addr, p.password)
	defer c.Close()

	cmd := fmt.Sprintf("kill %s", player)
	result, err := c.SendCommand(cmd)
	if err != nil {
		return err.Error()
	}

	return result
}

// NewPlayerService creates a new instance of the PlayerService
func NewPlayerService(addr string, password string) *PlayerService {
	return &PlayerService{addr: addr, password: password}
}

// Meetup joins player1 onto player2's position in the world
func (p *PlayerService) Meetup(ctx context.Context, player1, player2 string) string {
	c := rcon.NewConnection(ctx, p.addr, p.password)
	defer c.Close()

	cmd := fmt.Sprintf("tp %s %s", player1, player2)
	result, err := c.SendCommand(cmd)
	if err != nil {
		return err.Error()
	}

	return result
}

// Teleport moves a player to the 3D coordinates provided
func (p *PlayerService) Teleport(ctx context.Context, player, x, y, z string) string {
	c := rcon.NewConnection(ctx, p.addr, p.password)
	defer c.Close()

	cmd := fmt.Sprintf("tp %s %s %s %s", player, x, y, z)
	result, err := c.SendCommand(cmd)
	if err != nil {
		return err.Error()
	}

	return result
}

// Give adds the specified quantity of items into a player's inventory
func (p *PlayerService) Give(ctx context.Context, name string, item items.Item, quantity int) (string, error) {
	c := rcon.NewConnection(ctx, p.addr, p.password)
	cmd := fmt.Sprintf("give %s %s %d", name, item, quantity)
	result, err := c.SendCommand(cmd)
	if err != nil {
		return "", err
	}
	defer c.Close()

	return result, nil
}
