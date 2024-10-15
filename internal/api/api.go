package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type BitcoinFeeEstimate struct {
	SatPerVbyte float64 `json:"sat_per_vbyte"`
}

type BitcoinFeeEstimates map[string]BitcoinFeeEstimate

type BitcoinResponseData struct {
	Timestamp int64               `json:"timestamp"`
	Estimates BitcoinFeeEstimates `json:"estimates"`
}

func FetchBitcoinFee() (BitcoinResponseData, error) {
	resp, err := http.Get("https://bitcoiner.live/api/fees/estimates/latest")
	if err != nil {
		log.Printf("Error fetching Bitcoin fee: %v", err)
		return BitcoinResponseData{}, err
	}
	defer resp.Body.Close()

	var data BitcoinResponseData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Printf("Error decoding response: %v", err)
		return BitcoinResponseData{}, err
	}

	return data, nil
}
