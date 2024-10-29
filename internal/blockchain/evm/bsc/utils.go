package bsc

import "github.com/maindotmarcell/blockchain-fees/internal/blockchain/evm"

func ConvertWeiToBnb(wei float64) float64 {
	return evm.ConvertWeiToEth(wei)
}
