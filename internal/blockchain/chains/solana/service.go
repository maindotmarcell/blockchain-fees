package solana

import (
	"fmt"
)

type SolanaService struct {
	apiClient *APIClient
}

func NewSolanaService() *SolanaService {
	return &SolanaService{
		apiClient: NewSolanaAPIClient(),
	}
}

func (s *SolanaService) EstimateNetworkFee() (float64, error) {
	data, err := s.apiClient.FetchFee()
	if err != nil {
		return 0, fmt.Errorf("error fetching Solana fee: %v", err)
	}

	solanaResponseData, ok := data.(SolanaResponseData)
	if !ok {
		return 0, fmt.Errorf("invalid response data type")
	}

	lamportsPerSignature := solanaResponseData.Result.Value.FeeCalculator.LamportsPerSignature
	solFee := calculateFee(lamportsPerSignature)

	return solFee, nil
}

func calculateFee(fee uint64) float64 {
	lamportFee := fee * NUM_SIGNATURES
	return ConvertLamportsToSol(lamportFee)
}
