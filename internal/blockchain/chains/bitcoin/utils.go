package bitcoin

func convertSatoshisToBitcoin(satoshis float64) float64 {
	return satoshis / satoshisPerBitcoin
}
