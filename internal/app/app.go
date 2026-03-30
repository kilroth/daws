package app

import (
	"context"
	md "daws/internal/models"
	sm "daws/internal/storage"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
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

func (a *App) OpenFolderDialog() (string, error) {
	path, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title:            "Select Data Directory",
		DefaultDirectory: a.SM.GetDataDir(),
	})
	if err != nil {
		return "", md.Logger.Error("Failed to open folder dialog: %v", err)
	}
	return path, nil
}
