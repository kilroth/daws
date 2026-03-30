package models

type AppConfig struct {
	AWSAccessKey string `json:"awsAccessKey"`
	AWSSecretKey string `json:"awsSecretKey"`
	AWSRegion    string `json:"awsRegion"`
	Theme        string `json:"theme"`
	DataPath     string `json:"dataPath"`
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
