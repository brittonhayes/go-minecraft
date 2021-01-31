package go_minecraft

import (
	"context"
	"fmt"
	"github.com/brittonhayes/go-minecraft/internal/rcon"
	"strconv"
)

const (
	WeatherRain        Weather    = "rain"
	WeatherThunder     Weather    = "thunder"
	WeatherSnow        Weather    = "snow"
	ModeSurvival       Mode       = "survival"
	ModeCreative       Mode       = "creative"
	DifficultyEasy     Difficulty = "easy"
	DifficultyMedium   Difficulty = "normal"
	DifficultyHard     Difficulty = "hard"
	DifficultyPeaceful Difficulty = "peaceful"
)

var _ World = (*WorldService)(nil)

// World is the primary interface for the all world-related commands
type World interface {
	Weather(ctx context.Context, w Weather) string
	Time(ctx context.Context, time int) string
	Seed(ctx context.Context) string
	Mode(ctx context.Context, m Mode) string
	Difficulty(ctx context.Context, d Difficulty) string
}

type (
	Weather      string
	Mode         string
	Difficulty   string
	WorldService struct {
		addr     string
		password string
	}
)

func NewWorldService(addr string, password string) *WorldService {
	return &WorldService{addr: addr, password: password}
}

func (ws *WorldService) Difficulty(ctx context.Context, d Difficulty) string {
	c := rcon.NewConnection(ctx, ws.addr, ws.password)
	defer c.Close()

	cmd := fmt.Sprintf("difficulty %s", string(d))
	result, err := c.SendCommand(cmd)
	if err != nil {
		return err.Error()
	}

	return result
}

func (ws *WorldService) Mode(ctx context.Context, m Mode) string {
	c := rcon.NewConnection(ctx, ws.addr, ws.password)
	defer c.Close()

	cmd := fmt.Sprintf("gamemode %s", string(m))
	result, err := c.SendCommand(cmd)
	if err != nil {
		return err.Error()
	}

	return result
}

// Seed returns the world seed
func (ws *WorldService) Seed(ctx context.Context) string {
	c := rcon.NewConnection(ctx, ws.addr, ws.password)
	defer c.Close()

	result, err := c.SendCommand("seed")
	if err != nil {
		return err.Error()
	}

	return result
}

// Weather sets the weather in the world
func (ws *WorldService) Weather(ctx context.Context, w Weather) string {
	c := rcon.NewConnection(ctx, ws.addr, ws.password)
	defer c.Close()

	cmd := fmt.Sprintf("weather %s", string(w))
	result, err := c.SendCommand(cmd)
	if err != nil {
		return err.Error()
	}

	return result
}

// Time sets the time in the world
func (ws *WorldService) Time(ctx context.Context, time int) string {
	c := rcon.NewConnection(ctx, ws.addr, ws.password)
	defer c.Close()

	cmd := fmt.Sprintf("weather %s", strconv.Itoa(time))
	result, err := c.SendCommand(cmd)
	if err != nil {
		return err.Error()
	}

	return result
}
