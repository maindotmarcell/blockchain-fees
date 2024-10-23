package blockchain

type Blockchain interface {
	Name() string
	CurrencySymbol() string
	GetFee() (fee float64, err error)
}

type Service interface {
	EstimateNetworkFee() (fee float64, err error)
	CalculateFee(amount float64) (fee float64, err error)
}

type APIClient interface {
	FetchFee() (responseData interface{}, err error)
}
