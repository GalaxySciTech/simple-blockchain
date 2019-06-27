package core

import (
	"crypto/sha256"
	"encoding/hex"
)

func CalculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + block.PrevHash
	s := sha256.New()
	s.Write([]byte(record))
	hashed := s.Sum(nil)
	return hex.EncodeToString(hashed)
}
