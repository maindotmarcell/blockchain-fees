package bitcoin

import (
	"fmt"

	"github.com/maindotmarcell/blockchain-fees/internal/blockchain"
)

type BitcoinService struct {
	apiClient blockchain.APIClient
}

func NewBitcoinService() *BitcoinService {
	return &BitcoinService{
		apiClient: NewBitcoinAPIClient(),
	}
}

func (s *BitcoinService) EstimateNetworkFee() (float64, error) {
	data, err := s.apiClient.FetchFee()
	if err != nil {
		return 0, fmt.Errorf("error fetching Bitcoin fee: %v", err)
	}

	bitcoinResponseData, ok := data.(BitcoinResponseData)
	if !ok {
		return 0, fmt.Errorf("invalid response data type")
	}

	satPerVbyte := bitcoinResponseData.Estimates[defaultFeeEstimate].SatPerVbyte
	satFeeForTargetTransaction := satPerVbyte * float64(targetTransactionSize)
	bitcoinFee, err := s.calculateFee(satFeeForTargetTransaction)
	if err != nil {
		return 0, fmt.Errorf("error calculating fee: %v", err)
	}

	return bitcoinFee, nil
}

func (s *BitcoinService) calculateFee(satPerVbyte float64) (float64, error) {
	satFeeForTargetTransaction := satPerVbyte * float64(targetTransactionSize)
	bitcoinFee := convertSatoshisToBitcoin(satFeeForTargetTransaction)

	return bitcoinFee, nil
}
