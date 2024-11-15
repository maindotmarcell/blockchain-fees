package ethereum

import "github.com/maindotmarcell/blockchain-fees/internal/blockchain"

type Ethereum struct {
	service blockchain.Service
}

func New() *Ethereum {
	return &Ethereum{
		service: NewEthService(),
	}
}

func (e *Ethereum) Name() string {
	return "Ethereum"
}

func (e *Ethereum) CurrencySymbol() string {
	return "ETH"
}

func (e *Ethereum) GetFee() (float64, error) {
	return e.service.EstimateNetworkFee()
}
