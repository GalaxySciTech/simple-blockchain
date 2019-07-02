package main

import (
	"simple-blockchain/config"
	"simple-blockchain/core"
	"simple-blockchain/net"
)

func main() {

	core.Start()
	config.Init()
	net.Start()

}
