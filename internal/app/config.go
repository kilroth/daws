package app

import (
	md "daws/internal/models"
	sm "daws/internal/storage"
	"daws/internal/utils"
)

func (a *App) ConfigSave(config md.AppConfig) error {
	err := config.AWSCredentials.EncryptAll()
	if err != nil {
		return utils.Logger.Error("Failed to encrypt AWS credentials: %v", err)
	}

	return a.SM.SaveConfig(config)
}

func (a *App) ConfigLoad() (md.AppConfig, error) {
	storage := sm.NewStorageManager("")
	var config md.AppConfig
	var err error
	err = storage.LoadConfig(&config)
	if err != nil {
		config = md.AppConfig{
			Theme: "light",
		}
		utils.Logger.Info("No existing config found, creating new one with defaults")
		// create new storage manager
		a.SM = sm.NewStorageManager(config.DataPath)
		// manager sets default data path
		config.DataPath = a.SM.GetDataDir()
		// save that into our config.json
		storage.SaveConfig(config)
		return config, nil
	}
	a.SM = sm.NewStorageManager(config.DataPath)
	return config, nil
}

// get decrypted AWS access key for use in the app
func (a *App) GetDefaultAWSAccessKey() (string, error) {
	return a.Config.AWSCredentials.GetDecryptedAccessKey()
}

// get decrypted AWS secret key for use in the app
func (a *App) GetDefaultAWSSecretKey() (string, error) {
	return a.Config.AWSCredentials.GetDecryptedSecretKey()
}
