package app

import (
	md "daws/internal/models"
	sm "daws/internal/storage"
)

func (a *App) ConfigSave(config *md.AppConfig) error {
	return a.SM.SaveConfig(*config)
}

func (a *App) ConfigLoad() (md.AppConfig, error) {
	storage := sm.NewStorageManager("")
	var config md.AppConfig
	var err error
	err = storage.LoadConfig(&config)
	if err != nil {
		config = md.AppConfig{
			AWSAccount: "",
			AWSToken:   "",
			Theme:      "light",
			DataPath:   "",
		}
		storage.SaveConfig(config)
	}
	a.SM = sm.NewStorageManager(config.DataPath)
	return config, err
}
