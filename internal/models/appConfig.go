package models

type AppConfig struct {
	AWSCredentials AWSCredentials `json:"awsCredentials"`
	Theme          string         `json:"theme"`
	DataPath       string         `json:"dataPath"`
}

func (c *AppConfig) Save() error {
	var err error
	// write to file
	return err
}

func (c *AppConfig) Load() (*AppConfig, error) {
	var err error
	return c, err
}
