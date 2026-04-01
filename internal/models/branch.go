package models

import "daws/internal/utils"

type Branch struct {
	Name           string         `json:"name"`
	Slug           string         `json:"slug"`
	GitName        string         `json:"gitName"`
	Description    string         `json:"description"`
	AWSCredentials AWSCredentials `json:"awsCredentials"`
	AWSLocations   []AWSLocation  `json:"awsLocations"`
	IsActive       bool           `json:"isActive"`
	Archived       bool           `json:"archived"`
	TimeData       TimeData       `json:"timeData"`
}

func (b *Branch) Deploy() error {
	var err error
	return err
}

func (b *Branch) RunTests() error {
	var err error
	return err
}

func (b *Branch) Save_Preprocess() error {
	b.TimeData.Create()
	b.Slug = utils.Slugify(b.Name)
	return b.AWSCredentials.EncryptAll()
}
