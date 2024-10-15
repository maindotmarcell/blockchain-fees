package main

import (
	"github.com/maindotmarcell/blockchain-fees/internal/app"
)

func main() {
	app := app.New()
	app.Run()
}
