package bsc

import (
	"fmt"
	"strconv"

	"github.com/maindotmarcell/blockchain-fees/internal/blockchain"
)

type BscService struct {
	apiClient blockchain.APIClient
}

func NewBscService() *BscService {
	return &BscService{
		apiClient: NewBscAPIClient(),
	}
}

func (s *BscService) EstimateNetworkFee() (float64, error) {
	data, err := s.apiClient.FetchFee()
	if err != nil {
		return 0, fmt.Errorf("error fetching Bsc fee: %v", err)
	}

	bscResponseData, ok := data.(BscResponseData)
	if !ok {
		return 0, fmt.Errorf("invalid response data type")
	}
	// fmt.Println(bscResponseData) // DEBUG

	hexFeePerGas := bscResponseData.Result
	weiFee, err := s.calculateFee(hexFeePerGas)
	if err != nil {
		return 0, fmt.Errorf("error calculating fee: %v", err)
	}

	return ConvertWeiToBnb(weiFee), nil
}

func (s *BscService) calculateFee(hexFeePerGas string) (float64, error) {
	gasLimit := GAS_LIMIT
	feePerGas, err := strconv.ParseInt(hexFeePerGas[2:], 16, 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing priority fee: %v", err)
	}
	weiFee := float64(feePerGas) * float64(gasLimit)

	return weiFee, nil
}
