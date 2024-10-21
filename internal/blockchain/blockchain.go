package blockchain

type Blockchain interface {
	Name() string
	GetFee() (fee float64, err error)
	FetchApi() (responseData interface{}, err error)
	CalculateFee(satPerVbyte float64) (fee float64)
}
