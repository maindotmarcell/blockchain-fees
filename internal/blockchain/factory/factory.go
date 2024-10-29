package factory

import (
	"github.com/maindotmarcell/blockchain-fees/internal/blockchain"
	"github.com/maindotmarcell/blockchain-fees/internal/blockchain/bitcoin"
	"github.com/maindotmarcell/blockchain-fees/internal/blockchain/evm/bsc"
	"github.com/maindotmarcell/blockchain-fees/internal/blockchain/evm/ethereum"
)

func New(blockchainName string) blockchain.Blockchain {
	switch blockchainName {
	case "bitcoin":
		return bitcoin.New()
	case "ethereum":
		return ethereum.New()
	case "bsc":
		return bsc.New()
	default:
		return nil
	}
}
