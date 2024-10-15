package app

import "fmt"

type App struct{}

func New() *App {
	return &App{}
}

func (*App) FetchFees() {
	fmt.Println("Fetching fees...")
}
