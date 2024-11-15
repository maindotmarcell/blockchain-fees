package factory

import (
	"github.com/maindotmarcell/blockchain-fees/internal/blockchain"
	"github.com/maindotmarcell/blockchain-fees/internal/blockchain/chains/bitcoin"
	"github.com/maindotmarcell/blockchain-fees/internal/blockchain/chains/bsc"
	"github.com/maindotmarcell/blockchain-fees/internal/blockchain/chains/ethereum"
	"github.com/maindotmarcell/blockchain-fees/internal/blockchain/chains/solana"
)

func New(blockchainName string) blockchain.Blockchain {
	switch blockchainName {
	case "bitcoin":
		return bitcoin.New()
	case "ethereum":
		return ethereum.New()
	case "bsc":
		return bsc.New()
	case "solana":
		return solana.New()
	default:
		return nil
	}
}
