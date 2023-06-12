package config

type MinIO struct {
	Endpoint        string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	AccessKeyID     string `mapstructure:"accessKeyID" json:"accessKeyID" yaml:"accessKeyID"`
	SecretAccessKey string `mapstructure:"secretAccessKey" json:"secretAccessKey" yaml:"secretAccessKey"`
}
