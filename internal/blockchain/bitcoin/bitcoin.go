package bitcoin

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Bitcoin struct {
	apiUrl string
}

func New(apiUrl string) *Bitcoin {
	return &Bitcoin{apiUrl: apiUrl}
}

func (b *Bitcoin) Name() string {
	return "Bitcoin"
}

func (b *Bitcoin) GetFee() (float64, error) {
	data, err := b.FetchApi()
	if err != nil {
		return 0, fmt.Errorf("error fetching Bitcoin fee: %v", err)
	}

	bitcoinResponseData, ok := data.(BitcoinResponseData)
	if !ok {
		return 0, fmt.Errorf("invalid response data type")
	}

	satPerVbyte := bitcoinResponseData.Estimates[defaultFeeEstimate].SatPerVbyte
	satFeeForTargetTransaction := satPerVbyte * float64(targetTransactionSize)
	bitcoinFee := b.CalculateFee(satFeeForTargetTransaction)

	return bitcoinFee, nil
}

func (b *Bitcoin) FetchApi() (interface{}, error) {
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

func (b *Bitcoin) CalculateFee(satPerVbyte float64) float64 {
	satFeeForTargetTransaction := satPerVbyte * float64(targetTransactionSize)
	bitcoinFee := satFeeForTargetTransaction / satoshisPerBitcoin
	return bitcoinFee
}
