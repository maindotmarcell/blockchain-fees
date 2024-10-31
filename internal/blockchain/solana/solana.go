package solana

import "github.com/maindotmarcell/blockchain-fees/internal/blockchain"

type Solana struct {
	service blockchain.Service
}

func New() *Solana {
	return &Solana{
		service: NewSolanaService(),
	}
}

func (s *Solana) Name() string {
	return "Solana"
}

func (s *Solana) CurrencySymbol() string {
	return "SOL"
}

func (s *Solana) GetFee() (float64, error) {
	return s.service.EstimateNetworkFee()
}
