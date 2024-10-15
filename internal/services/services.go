package services

import (
	"log"

	"github.com/maindotmarcell/blockchain-fees/internal/api"
)

func GetBitcoinFee() float64 {
	bitcoinResponseData, err := api.FetchBitcoinFee()
	if err != nil {
		log.Printf("Error fetching Bitcoin fee: %v", err)
		return 0
	}
	satPerVbyte := bitcoinResponseData.Estimates["30"].SatPerVbyte
	targetTransactionSize := 140
	satFeeForTargetTransaction := satPerVbyte * float64(targetTransactionSize)
	bitcoinFee := satFeeForTargetTransaction / 100000000

	return bitcoinFee
}
