package main

import (
	"fmt"
	"simple-blockchain/core"
	"time"
)

func main() {
	blo := core.Block{Index: 1, Timestamp: time.Now().String(), PrevHash: "3c5101b64f2d810b579159de73c40ec40124b05f01a31053e8b8d2618ffc56b3"}
	hash:= core.CalculateHash(blo)
	blo.Hash=hash
	fmt.Println(blo)
	for{
		newBlock,_:=core.GenerateBlock(blo)
		fmt.Println(newBlock)
		if(core.IsBlockValid(newBlock,blo)){
			blo=newBlock

		}else {
			fmt.Println("block generate failed reason: new block verification failed ")
			return
		}
		time.Sleep(time.Second)
	}

}
