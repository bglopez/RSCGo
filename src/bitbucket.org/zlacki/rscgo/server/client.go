package server

import (
	"bitbucket.org/zlacki/rscgo/entity"
	"fmt"
	"net"
	"strconv"
)

type client struct {
	channel
	cipherKey int64
	uID       uint8
	ip        string
	index     int
	kill      chan struct{}
	player    *entity.Player
}

//unregister Clean up resources and unregister the receiver from the global clientList.
func (c *client) unregister() {
	close(c.kill)
	close(c.send)
	fmt.Println("Unregistering client" + c.String())
	if err := c.socket.Close(); err != nil {
		fmt.Printf("WARNING: Error closing listener for client%s\n", c.String())
		fmt.Println(err)
	}
	activeClients.remove(c.index)
}

//startWriter todo: do I need this?
func (c *client) startWriter() {
	go func() {
		defer func() {
			c.unregister()
		}()
		for {
			select {
			case p := <-c.send:
				if p != nil {
					c.writePacket(p)
				}
				continue
			case <-c.kill:
				return;
			}
		}
	}()
}

//startReader Creates a new goroutine to handle all incoming network events for the receiver client.
// This goroutine will also automatically handle cleanup for client disconnections, and handle incoming I/O errors
// and disconnect the related client appropriately.
func (c *client) startReader() {
	go func() {
		for {
			select {
			case <-c.kill:
				return
			default:
				p, err := c.readPacket()
				if p == nil {
					return
				}
				if err != nil {
					fmt.Println(err.Error())
					if err.ping || err.closed {
						return
					}
				}
				c.handlePacket(p)
			}
		}
	}()
}

//newClient Creates a new instance of a client, registers it with the global clientList, and returns it.
func newClient(socket net.Conn) *client {
	c := &client{channel: channel{socket: socket, send: make(chan *packet)}, cipherKey: -1, ip: getIPFromConn(socket), index: -1, kill: make(chan struct{}, 1), player: entity.NewPlayer()}
	c.startReader()
	c.startWriter()
	return c
}

//String Returns a string populated with some of the more identifying fields from the receiver client.
func (c *client) String() string {
	return "{idx:'" + strconv.Itoa(c.index) + "', ip:'" + c.ip + "'};"
}
