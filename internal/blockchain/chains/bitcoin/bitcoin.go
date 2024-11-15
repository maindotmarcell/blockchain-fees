package bitcoin

import "github.com/maindotmarcell/blockchain-fees/internal/blockchain"

type Bitcoin struct {
	service blockchain.Service
}

func New() *Bitcoin {
	return &Bitcoin{
		service: NewBitcoinService(),
	}
}

func (b *Bitcoin) Name() string {
	return "Bitcoin"
}

func (b *Bitcoin) CurrencySymbol() string {
	return "BTC"
}

func (b *Bitcoin) GetFee() (float64, error) {
	return b.service.EstimateNetworkFee()
}
