package entity

import (
	"meChain2.0/src/blockchain/utils"
)

type Blockchain_Iterator struct {
	CurrentHash []byte
	Chain       *Blockchain
}

func (bc *Blockchain) Iterator() *Blockchain_Iterator {
	return &Blockchain_Iterator{CurrentHash: bc.GetTopBlock().Hash, Chain: bc}
}

func (iterator *Blockchain_Iterator) Next() *Block {
	block := iterator.Chain.GetBlock(iterator.CurrentHash)
	iterator.CurrentHash = block.PrevBlockHash
	return block
}

func (bc *Blockchain) GetBlock(hash []byte) *Block {
	for _, block := range bc.Blocks {
		if utils.IsEqual(block.Hash, hash) {
			return block
		}
	}
	return nil
}

func (iterator *Blockchain_Iterator) PrintChainIterator() {
	for {
		block := iterator.Next()
		block.PrintBlock()
		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}
