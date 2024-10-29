package ethereum

import "github.com/maindotmarcell/blockchain-fees/internal/blockchain/evm"

func ConvertWeiToEth(weiFee float64) float64 {
	return evm.ConvertWeiToEth(weiFee)
}
