package client

import (
	"fmt"
	"l1/21/adapter"
)

type Client struct{}

func (c *Client) InsertLightningConnectorIntoComputer(com adapter.Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}
