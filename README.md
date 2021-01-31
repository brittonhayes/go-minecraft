# Go Minecraft

[![Go Reference](https://pkg.go.dev/badge/github.com/brittonhayes/go-minecraft.svg)](https://pkg.go.dev/github.com/brittonhayes/go-minecraft)
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/brittonhayes/go-minecraft?color=blue&label=Latest%20Version&sort=semver)
[![Go Report Card](https://goreportcard.com/badge/github.com/brittonhayes/go-minecraft)](https://goreportcard.com/report/github.com/brittonhayes/go-minecraft)

<img width="300" src="https://raw.githubusercontent.com/egonelbre/gophers/10cc13c5e29555ec23f689dc985c157a8d4692ab/vector/fairy-tale/knight.svg"></img>

> A Go Client for the Minecraft RCON protocol

## Installation

Install with the go get command

```shell
go get github.com/brittonhayes/go-minecraft
```

## Documentation

View the full docs on [pkg.go.dev](https://pkg.go.dev/github.com/brittonhayes/go-minecraft)

## Usage

Using the package is as easy as create client, pick the endpoint, and run the method. This applies across every data
type, so it is consistent across the board. Here's a simple example of how to give a player an item in-game.

```go
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
```

## Examples

For example uses of the package, check out the [example](./example) directory

## Development

If you'd like to contribute to go-minecraft, make sure you have mage installed: https://magefile.org

```shell
# Download dependencies and run tests
mage download
mage test
```
