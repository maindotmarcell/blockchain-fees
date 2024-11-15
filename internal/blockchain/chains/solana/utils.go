package solana

func ConvertLamportsToSol(lamports uint64) float64 {
	return float64(lamports) / 1e9
}
