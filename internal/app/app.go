package app

import (
	"fmt"
	"sync"
	"time"

	"github.com/maindotmarcell/blockchain-fees/internal/services"
)

type App struct{}

func New() *App {
	return &App{}
}

func (*App) FetchFees() {
	datetime := time.Now().Format("2006-01-02 15:04:05")

	var wg sync.WaitGroup
	var bitcoinFee float64

	wg.Add(1)
	go func() {
		defer wg.Done()
		bitcoinFee = services.GetBitcoinFee()
	}()

	wg.Wait()
	fmt.Printf("Fee for Bitcoin at %v: %.8f BTC\n", datetime, bitcoinFee)
}

func (*App) Run() {
	app := New()
	app.FetchFees()
}
