package bsc

import "github.com/maindotmarcell/blockchain-fees/internal/blockchain"

func ConvertWeiToBnb(wei float64) float64 {
	return blockchain.ConvertWeiToEth(wei)
}
