package blockchain

type Blockchain interface {
	GetFee() (float64, error)
	fetchApi() (interface{}, error)
	Name() string
}
