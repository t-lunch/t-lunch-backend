package config

import "github.com/spf13/viper"

type Config struct {
	Database struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"dbname"`
		SSLMode  string `mapstructure:"sslmode"`
	} `mapstructure:"database"`

	Server struct {
		HTTPPort int `mapstructure:"http_port"`
	} `mapstructure:"server"`

	Jwt struct {
		Secret            string `mapstructure:"secret"`
		AccessExpiration  int    `mapstructure:"access_expiration"`
		RefreshExpiration int    `mapstructure:"refresh_expiration"`
	} `mapstructure:"jwt"`
}

func NewConfig(configName string) (*Config, error) {
	viper.AddConfigPath("./configs")
	viper.SetConfigType("yaml")
	viper.SetConfigName(configName)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
