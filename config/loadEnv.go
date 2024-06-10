package config

import "github.com/spf13/viper"

type DBConfig struct {
	Username string `mapstructure:"USERNAME_DB"`
	Password string `mapstructure:"PASSWORD_DB"`
	DB_name  string `mapstructure:"DB_NAME"`
	IP       string `mapstructure:"IP"`
	Port     string `mapstructure:"PORT"`
}

func LoadEnv(path string) (dbConfig DBConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&dbConfig)
	return
}
