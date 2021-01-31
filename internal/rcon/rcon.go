package rcon

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"io"
	"net"
)

const (
	BadAuth            = -1
	PayloadMaxSize     = 1460
	ErrBadAuth         = "failed to authenticate"
	ErrPayloadTooLarge = "payload is too large"
)

const (
	PacketResponse = iota
	_
	PacketCommand
	PacketLogin
)

type Connection struct {
	conn     net.Conn
	password string
}

type packetType int32

type Header struct {
	Size      int32
	RequestID int32
	Type      packetType
}

type Connecter interface {
	Open(addr string, password string)
	SendCommand(command string) (string, error)
	Close()
}

func NewConnection(ctx context.Context, addr string, password string) *Connection {
	c := &Connection{}
	err := c.open(ctx, addr, password)
	if err != nil {
		panic(err)
	}

	return c
}

func (c *Connection) open(ctx context.Context, addr string, password string) error {
	var d net.Dialer
	conn, err := d.DialContext(ctx, "tcp", addr)
	if err != nil {
		return err
	}

	c.conn = conn
	c.password = password
	return nil
}

func (c *Connection) Close() {
	_ = c.conn.Close()
}

// SendCommand sends a command to the server and returns the result (often nothing).
func (c *Connection) SendCommand(command string) (string, error) {
	// Send the packet.
	if len([]byte(command)) > PayloadMaxSize {
		return "", errors.New(ErrPayloadTooLarge)
	}
	head, payload, err := c.sendPacket(PacketCommand, []byte(command))
	if err != nil {
		return "", err
	}

	// Auth was bad, throw error.
	if head.RequestID == BadAuth {
		return "", errors.New(ErrBadAuth)
	}
	return string(payload), nil
}

// Authenticate authenticates the user with the server.
func (c *Connection) Authenticate() error {
	// Send the packet.
	head, _, err := c.sendPacket(PacketLogin, []byte(c.password))
	if err != nil {
		return err
	}

	// If the credentials were bad, throw error.
	if head.RequestID == BadAuth {
		return errors.New(ErrBadAuth)
	}

	return nil
}

// sendPacket sends the binary packet representation to the server and returns the response.
func (c *Connection) sendPacket(t packetType, p []byte) (Header, []byte, error) {
	// Generate the binary packet.
	packet, err := encodePacket(t, p)
	if err != nil {
		return Header{}, nil, err
	}

	// Send the packet over the wire.
	_, err = c.conn.Write(packet)
	if err != nil {
		return Header{}, nil, err
	}
	// Receive and decode the response.
	head, payload, err := decodePacket(c.conn)
	if err != nil {
		return Header{}, nil, err
	}

	return head, payload, nil
}

// encodePacket encodes the packet type and payload into a binary representation to send over the wire.
func encodePacket(t packetType, p []byte) ([]byte, error) {
	// Generate a random request ID.
	pad := [2]byte{}
	length := int32(len(p) + 10)
	var buf bytes.Buffer
	_ = binary.Write(&buf, binary.LittleEndian, length)
	_ = binary.Write(&buf, binary.LittleEndian, int32(0))
	_ = binary.Write(&buf, binary.LittleEndian, t)
	_ = binary.Write(&buf, binary.LittleEndian, p)
	_ = binary.Write(&buf, binary.LittleEndian, pad)
	// Notchian server doesn't like big packets :(
	if buf.Len() >= 1460 {
		return nil, errors.New("packet too big when packetising")
	}
	// Return the bytes.
	return buf.Bytes(), nil
}

// decodePacket decodes the binary packet into a native Go struct.
func decodePacket(r io.Reader) (Header, []byte, error) {
	head := Header{}
	err := binary.Read(r, binary.LittleEndian, &head)
	if err != nil {
		return Header{}, nil, err
	}
	payload := make([]byte, head.Size-8)
	_, err = io.ReadFull(r, payload)
	if err != nil {
		return Header{}, nil, err
	}

	// Some basic sanity checking
	if head.Type != PacketResponse && head.Type != PacketCommand {
		return Header{}, nil, errors.New("bad packet type")
	}
	return head, payload[:len(payload)-2], nil
}
