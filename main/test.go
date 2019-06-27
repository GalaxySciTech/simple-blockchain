package main

import (
	"fmt"
	"simple-blockchain/core"
	"time"
)

func main() {
	b := core.Block{Index: 1, Timestamp: time.RFC3339, PrevHash: "3c5101b64f2d810b579159de73c40ec40124b05f01a31053e8b8d2618ffc56b3"}
	hash:= core.CalculateHash(b)
	b.Hash=hash
	fmt.Println(b)
}
