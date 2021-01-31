package go_minecraft

type Client struct {
	*PlayerService
	*WorldService
}

// NewClient creates a new instance of the Minecraft client
func NewClient(host string, password string) *Client {
	return &Client{
		PlayerService: NewPlayerService(host, password),
		WorldService:  NewWorldService(host, password),
	}
}
