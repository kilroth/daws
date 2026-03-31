package app

import (
	"context"
	md "daws/internal/models"
	"daws/internal/utils"
)

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx

	data, config, err := a.StartupData()
	if err != nil {
		utils.Logger.Error("Failed to load startup data: %v", err)
		return
	}
	a.Data = data
	a.Config = config
}

func (a *App) StartupData() (md.AppData, md.AppConfig, error) {
	var err error
	var data md.AppData
	var config md.AppConfig

	// load the config
	config, err = a.ConfigLoad()
	if err != nil {
		return data, config, utils.Logger.Error("Failed to load config: %v", err)
	}

	// The sm makes sure the projects path exists, load the projects
	data, err = a.DataLoad()
	if err != nil {
		return md.AppData{}, config, utils.Logger.Error("Failed to load projects: %v", err)
	}

	return data, config, nil
}
