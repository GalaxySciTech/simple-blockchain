package main

import (
	"fmt"
	"simple-blockchain/core"
	"time"
)

func main() {
	genesisBlock := core.Block{Index: 0, Timestamp: time.Now().String(), PrevHash: ""}
	hash:= core.CalculateHash(genesisBlock)
	genesisBlock.Hash=hash
	blo:=genesisBlock
	core.BlockChain = append(core.BlockChain, genesisBlock)
	go func() {
		for{
			newBlock,_:=core.GenerateBlock(blo)
			if core.IsBlockValid(newBlock,blo) {
				blo=newBlock
				core.BlockChain= append(core.BlockChain, blo)
			}else {
				fmt.Println("block generate failed reason: new block verification failed ")
				return
			}
			time.Sleep(time.Second)
		}
	}()

}
