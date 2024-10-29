package ethereum

import (
	"fmt"
	"strconv"

	"github.com/maindotmarcell/blockchain-fees/internal/blockchain"
)

type EthService struct {
	apiClient blockchain.APIClient
}

func NewEthService() *EthService {
	return &EthService{
		apiClient: NewEthApiClient(),
	}
}

func (s *EthService) EstimateNetworkFee() (float64, error) {
	data, err := s.apiClient.FetchFee()
	if err != nil {
		return 0, fmt.Errorf("error fetching Bsc fee: %v", err)
	}

	ethResponseData, ok := data.(EthResponseData)
	if !ok {
		return 0, fmt.Errorf("invalid response data type")
	}
	// fmt.Println("ethResponseData:") // DEBUG
	// fmt.Println(ethResponseData)    // DEBUG

	hexFeePerGas := ethResponseData.Result
	weiFee, err := s.calculateFee(hexFeePerGas)
	if err != nil {
		return 0, fmt.Errorf("error calculating fee: %v", err)
	}

	return ConvertWeiToEth(weiFee), nil
}

func (s *EthService) calculateFee(hexFeePerGas string) (float64, error) {
	gasLimit := GAS_LIMIT
	feePerGas, err := strconv.ParseInt(hexFeePerGas[2:], 16, 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing priority fee: %v", err)
	}

	return float64(feePerGas) * float64(gasLimit), nil
}
