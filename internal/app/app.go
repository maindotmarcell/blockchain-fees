package app

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/maindotmarcell/blockchain-fees/internal/blockchain"
	"github.com/maindotmarcell/blockchain-fees/internal/blockchain/factory"
)

type App struct {
	blockchains []blockchain.Blockchain
}

func New() *App {
	return &App{
		blockchains: []blockchain.Blockchain{
			factory.New("bitcoin"),
			factory.New("bsc"),
			factory.New("ethereum"),
			factory.New("solana"),
		},
	}
}

func (app *App) GetFees() {
	datetime := time.Now().Format("2006-01-02 15:04:05")

	var wg sync.WaitGroup
	fees := make(map[string]Fee)

	for _, bc := range app.blockchains {
		wg.Add(1)
		go func(blockchain blockchain.Blockchain) {
			defer wg.Done()
			blockchainFee, err := blockchain.GetFee()
			if err != nil {
				log.Printf("Error fetching fee for %v: %v", blockchain.Name(), err)
			}
			fees[blockchain.Name()] = Fee{
				Blockchain:     blockchain.Name(),
				Fee:            blockchainFee,
				CurrencySymbol: blockchain.CurrencySymbol(),
			}
		}(bc)
	}

	wg.Wait()

	for _, fee := range fees {
		fmt.Printf("Fee for %v at %v: %.8f %v\n", fee.Blockchain, datetime, fee.Fee, fee.CurrencySymbol)
	}

}

func (app *App) Run() {
	app.GetFees()
}
