package factory

import (
	"github.com/maindotmarcell/blockchain-fees/internal/blockchain"
	"github.com/maindotmarcell/blockchain-fees/internal/blockchain/bitcoin"
	"github.com/maindotmarcell/blockchain-fees/internal/blockchain/bsc"
)

func New(blockchainName string) blockchain.Blockchain {
	switch blockchainName {
	case "bitcoin":
		return bitcoin.New()
	case "bsc":
		return bsc.New()
	default:
		return nil
	}
}
