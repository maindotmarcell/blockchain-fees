package ethereum

import "github.com/maindotmarcell/blockchain-fees/internal/blockchain"

func ConvertWeiToEth(weiFee float64) float64 {
	return blockchain.ConvertWeiToEth(weiFee)
}
