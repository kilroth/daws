package models

type Branch struct {
	Name         string        `json:"name"`
	GitName      string        `json:"gitName"`
	Description  string        `json:"description"`
	AWSAccount   string        `json:"awsAccount"`
	AWSLocations []AWSLocation `json:"awsLocations"`
	LastBuild    string        `json:"lastBuild"`
	IsActive     bool          `json:"isActive"`
	Archived     bool          `json:"archived"`
}

func (b *Branch) Deploy() error {
	var err error
	return err
}

func (b *Branch) RunTests() error {
	var err error
	return err
}
