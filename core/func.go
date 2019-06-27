package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

func CalculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + block.PrevHash
	s := sha256.New()
	s.Write([]byte(record))
	hashed := s.Sum(nil)
	return hex.EncodeToString(hashed)
}


func GenerateBlock(oldBlock Block) (Block, error) {
	var newBlock Block
	t := time.Now()
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = CalculateHash(newBlock)

	return newBlock, nil
}

func IsBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}
	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}
	if CalculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}
