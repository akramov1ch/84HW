package config

import (
    "github.com/spf13/viper"
)

func LoadConfig() {
    viper.AddConfigPath(".")
    viper.SetConfigName("config")
    viper.SetConfigType("env")

    if err := viper.ReadInConfig(); err != nil {
        panic(err)
    }
}
