package app

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/maindotmarcell/blockchain-fees/internal/blockchain"
	"github.com/maindotmarcell/blockchain-fees/internal/blockchain/bitcoin"
)

type App struct {
	blockchains []blockchain.Blockchain
}

func New() *App {
	bitcoin := bitcoin.New("https://bitcoiner.live/api/fees/estimates/latest")
	return &App{
		blockchains: []blockchain.Blockchain{
			bitcoin,
		},
	}
}

func (app *App) FetchFees() {
	datetime := time.Now().Format("2006-01-02 15:04:05")

	var wg sync.WaitGroup
	fees := make(map[string]float64)

	for _, bc := range app.blockchains {
		wg.Add(1)
		go func(blockchain blockchain.Blockchain) {
			defer wg.Done()
			blockchainFee, err := blockchain.GetFee()
			if err != nil {
				log.Printf("Error fetching fee for %v: %v", blockchain.Name(), err)
			}
			fees[blockchain.Name()] = blockchainFee
		}(bc)
	}

	wg.Wait()

	for blockchain, fee := range fees {
		fmt.Printf("Fee for %v at %v: %.8f BTC\n", blockchain, datetime, fee)
	}

}

func (*App) Run() {
	app := New()
	app.FetchFees()
}
