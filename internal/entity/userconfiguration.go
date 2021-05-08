package entity

type UserConfig struct {
	Database struct{
		Driver string `yaml:"driver"`
		DSN string `yaml:"dsn"`
	} `yaml:"database"`
}
