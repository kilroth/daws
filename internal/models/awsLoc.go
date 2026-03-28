package models

type AWSLocation struct {
	Description  string       `json:"description"`
	Region       string       `json:"region"`
	ECRConfig    ECRConfig    `json:"ecr"`
	LambdaConfig LambdaConfig `json:"lambdaDistro"`
	S3Config     S3Config     `json:"s3Config"`
	Invalidate   bool         `json:"invalidate"`
	OSTarget     string       `json:"osTarget"` //linux/amd64, windows/amd64, linux/arm64, etc.
}

type S3Config struct {
	BucketName string `json:"bucketName"`
	Key        string `json:"key"`
}

type ECRConfig struct {
	RepositoryName string `json:"repositoryName"`
	ImageTag       string `json:"imageTag"`
}

type LambdaConfig struct {
	FunctionName string `json:"functionName"`
	Distro       string `json:"distro"`
}
