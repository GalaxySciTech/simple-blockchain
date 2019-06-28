package main

import (
	"github.com/joho/godotenv"
	"log"
	"simple-blockchain/core"
	"simple-blockchain/net"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	genesisBlock := core.Block{Index: 0, Timestamp: time.Now().String(), PrevHash: ""}
	hash := core.CalculateHash(genesisBlock)
	genesisBlock.Hash = hash
	blo := genesisBlock
	core.BlockChain = append(core.BlockChain, genesisBlock)
	go func() {
		for {
			newBlock, _ := core.GenerateBlock(blo)
			if core.IsBlockValid(newBlock, blo) {
				blo = newBlock
				core.BlockChain = append(core.BlockChain, blo)
				log.Println("generate new block ",newBlock)
			} else {
				log.Println("block generate failed reason: new block verification failed ")
				return
			}
			time.Sleep(time.Second)
		}
	}()

	_ = net.Start()

}
