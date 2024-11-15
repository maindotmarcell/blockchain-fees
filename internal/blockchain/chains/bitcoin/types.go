package bitcoin

type BitcoinFeeEstimate struct {
	SatPerVbyte float64 `json:"sat_per_vbyte"`
}

type BitcoinFeeEstimates map[string]BitcoinFeeEstimate

type BitcoinResponseData struct {
	Timestamp int64               `json:"timestamp"`
	Estimates BitcoinFeeEstimates `json:"estimates"`
}
