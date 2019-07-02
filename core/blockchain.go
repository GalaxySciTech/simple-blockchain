package core

import "time"

// BlockChain 是一个 Block 指针数组
type BlockChain struct {
	Blocks []*Block
}

var Bc *BlockChain

// NewBlockChain 创建一个有创世块的链
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

// AddBlock 向链中加入一个新块
// data 在实际中就是交易
func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(prevBlock.index, data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}


func Start(){
	Bc=NewBlockChain()
	go func() {
		i:=1
		for ; ;  {
			Bc.AddBlock(string(i))
			i++
			time.Sleep(time.Second)
		}
	}()
}