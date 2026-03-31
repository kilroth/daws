package app

import (
	md "daws/internal/models"
)

func (a *App) GetAccessKey(awsCredentials md.AWSCredentials) (string, error) {
	return awsCredentials.GetDecryptedAccessKey()
}

func (a *App) GetSecretKey(awsCredentials md.AWSCredentials) (string, error) {
	return awsCredentials.GetDecryptedSecretKey()
}
