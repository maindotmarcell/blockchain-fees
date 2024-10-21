package blockchain

type Blockchain interface {
	Name() string                                    // returns the name of the blockchain
	GetFee() (fee float64, err error)                // returns the fee for the blockchain
	FetchApi() (responseData interface{}, err error) // fetches the data from the blockchain's API
	CalculateFee(satPerVbyte float64) (fee float64)  // calculates the fee for the blockchain
}
