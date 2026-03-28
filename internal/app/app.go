package app

import (
	"context"
	md "daws/internal/models"
	sm "daws/internal/storage"
	"fmt"
)

// App struct
type App struct {
	ctx    context.Context
	Config md.AppConfig
	Data   md.AppData
	SM     *sm.StorageManager
}

func NewApp() *App {
	return &App{}
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
