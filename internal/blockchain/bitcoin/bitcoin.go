package bitcoin

import "github.com/maindotmarcell/blockchain-fees/internal/blockchain"

type Bitcoin struct {
	service blockchain.Service
}

func New(apiURL string) *Bitcoin {
	return &Bitcoin{
		service: NewBitcoinService(apiURL),
	}
}

func (b *Bitcoin) Name() string {
	return "Bitcoin"
}

func (b *Bitcoin) GetFee() (float64, error) {
	return b.service.EstimateNetworkFee()
}
