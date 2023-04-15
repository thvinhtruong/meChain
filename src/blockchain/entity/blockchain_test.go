package entity

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestMineBlock(t *testing.T) {
	bc := NewBlockchain()
	require.NotNil(t, bc)

	block1 := bc.MineBlock("vincent")
	require.NotNil(t, block1)

	block2 := bc.MineBlock("vincent")
	require.NotNil(t, block2)

	// bc.PrintChain()
	// bc.PrintAllTransactions()

	start := time.Now()
	bc.Iterator().PrintChainIterator()

	elapsed := time.Since(start)

	fmt.Println("Time elapsed: ", elapsed)

	// test serialization and deserialization
	data := bc.Serialize()
	fmt.Printf("Serialized data: %x\n", data)

	bc2 := DeserializeBlockchain(data)
	bc2.Iterator().PrintChainIterator()
}
