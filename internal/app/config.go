package app

import (
	md "daws/internal/models"
	"daws/internal/storage"
	sm "daws/internal/storage"
)

func (a *App) ConfigSave(config md.AppConfig) error {
	encKey, err := storage.Encrypt(config.AWSAccessKey)
	if err != nil {
		return md.Logger.Error("Failed to get encryption key: %v", err)
	}

	encSecretKey, err := storage.Encrypt(config.AWSSecretKey)
	if err != nil {
		return md.Logger.Error("Failed to get encryption key: %v", err)
	}

	config.AWSAccessKey = encKey
	config.AWSSecretKey = encSecretKey

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
		md.Logger.Info("No existing config found, creating new one with defaults")
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
func (a *App) GetAWSAccessKey() (string, error) {
	if a.Config.AWSAccessKey == "" {
		return "", nil
	}
	decKey, err := storage.Decrypt(a.Config.AWSAccessKey)
	if err != nil {
		return "", md.Logger.Error("Failed to get decryption key: %v", err)
	}
	return decKey, nil
}

// get decrypted AWS secret key for use in the app
func (a *App) GetAWSSecretKey() (string, error) {
	if a.Config.AWSSecretKey == "" {
		return "", nil
	}
	decKey, err := storage.Decrypt(a.Config.AWSSecretKey)
	if err != nil {
		return "", md.Logger.Error("Failed to get decryption key: %v", err)
	}
	return decKey, nil
}
