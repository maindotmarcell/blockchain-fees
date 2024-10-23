package bsc

func ConvertWeiToBnb(wei float64) float64 {
	return ConvertWeiToEth(wei)
}

func ConvertWeiToEth(wei float64) float64 {
	return wei / 1e18
}
