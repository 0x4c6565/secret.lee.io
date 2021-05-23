package main

import "github.com/spf13/viper"

func InitConfig() error {
	viper.SetConfigName("config")
	viper.SetEnvPrefix("secret")
	viper.AutomaticEnv()
	setConfigDefaults()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	return nil
}

func setConfigDefaults() {
	viper.SetDefault("storage_redis_host", "")
	viper.SetDefault("storage_redis_password", "")
	viper.SetDefault("storage_redis_port", 6379)
	viper.SetDefault("storage_redis_db", 0)
}
