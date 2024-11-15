package bsc

import "github.com/maindotmarcell/blockchain-fees/internal/blockchain"

type Bsc struct {
	service blockchain.Service
}

func New() *Bsc {
	return &Bsc{
		service: NewBscService(),
	}
}

func (b *Bsc) Name() string {
	return "BSC"
}

func (b *Bsc) CurrencySymbol() string {
	return "BNB"
}

func (b *Bsc) GetFee() (float64, error) {
	return b.service.EstimateNetworkFee()
}
