package factory

import (
	"github.com/maindotmarcell/blockchain-fees/internal/blockchain"
	"github.com/maindotmarcell/blockchain-fees/internal/blockchain/bitcoin"
)

func New(blockchainName string) blockchain.Blockchain {
	switch blockchainName {
	case "bitcoin":
		return bitcoin.New(bitcoin.URL)
	default:
		return nil
	}
}
