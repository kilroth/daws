package models

import (
	"daws/internal/utils"
)

type AWSCredentials struct {
	AccessKey string   `json:"accessKey"`
	SecretKey string   `json:"secretKey"`
	Regions   []string `json:"regions"`
}

func (c *AWSCredentials) Validate() error {
	var err error
	if c.AccessKey == "" || c.SecretKey == "" || len(c.Regions) == 0 {
		err = utils.Logger.Error("AWS credentials are incomplete")
	}

	// remove empty regions
	var validRegions []string
	for _, region := range c.Regions {
		if region != "" {
			validRegions = append(validRegions, region)
		}
	}
	c.Regions = validRegions
	return err
}

func (c *AWSCredentials) GetDecryptedAccessKey() (string, error) {
	var err error
	if c.AccessKey == "" {
		return "", utils.Logger.Error("AWS credentials are incomplete")
	}
	docKey, err := utils.Decrypt(c.AccessKey)
	if err != nil {
		return "", utils.Logger.Error("Failed to get decryption key: %v", err)
	}
	c.AccessKey = docKey
	return docKey, nil
}

func (c *AWSCredentials) GetDecryptedSecretKey() (string, error) {
	var err error
	if c.SecretKey == "" {
		return "", utils.Logger.Error("AWS credentials are incomplete")
	}
	docKey, err := utils.Decrypt(c.SecretKey)
	if err != nil {
		return "", utils.Logger.Error("Failed to get decryption key: %v", err)
	}
	c.SecretKey = docKey
	return docKey, nil
}

func (c *AWSCredentials) EncryptAll() error {
	var err error
	c.AccessKey, err = utils.Encrypt(c.AccessKey)
	if err != nil {
		return utils.Logger.Error("Failed to encrypt access key: %v", err)
	}
	c.SecretKey, err = utils.Encrypt(c.SecretKey)
	if err != nil {
		return utils.Logger.Error("Failed to encrypt secret key: %v", err)
	}
	return nil
}
