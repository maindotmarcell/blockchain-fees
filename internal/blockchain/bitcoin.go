package blockchain

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Bitcoin struct {
	apiUrl string
}

type BitcoinFeeEstimate struct {
	SatPerVbyte float64 `json:"sat_per_vbyte"`
}

type BitcoinFeeEstimates map[string]BitcoinFeeEstimate

type BitcoinResponseData struct {
	Timestamp int64               `json:"timestamp"`
	Estimates BitcoinFeeEstimates `json:"estimates"`
}

func NewBitcoin(apiUrl string) *Bitcoin {
	return &Bitcoin{apiUrl: apiUrl}
}

func (b *Bitcoin) Name() string {
	return "Bitcoin"
}

func (b *Bitcoin) GetFee() (float64, error) {
	data, err := b.fetchApi()
	if err != nil {
		return 0, fmt.Errorf("error fetching Bitcoin fee: %v", err)
	}

	bitcoinResponseData, ok := data.(BitcoinResponseData)
	if !ok {
		return 0, fmt.Errorf("invalid response data type")
	}

	satPerVbyte := bitcoinResponseData.Estimates["30"].SatPerVbyte
	targetTransactionSize := 140
	satFeeForTargetTransaction := satPerVbyte * float64(targetTransactionSize)
	bitcoinFee := satFeeForTargetTransaction / 100000000

	return bitcoinFee, nil
}

func (b *Bitcoin) fetchApi() (interface{}, error) {
	resp, err := http.Get(b.apiUrl)
	if err != nil {
		return nil, fmt.Errorf("error fetching Bitcoin fee: %v", err)
	}
	defer resp.Body.Close()

	var data BitcoinResponseData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return data, nil
}
