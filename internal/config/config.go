package config

import (
    "github.com/spf13/viper"
)

type Config struct {
    MongoURI string
    DBName   string
    Port     string
}

func Load() (*Config, error) {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    viper.AutomaticEnv()

    err := viper.ReadInConfig()
    if err != nil {
        return nil, err
    }

    cfg := &Config{
        MongoURI: viper.GetString("MONGO_URI"),
        DBName:   viper.GetString("DB_NAME"),
        Port:     viper.GetString("PORT"),
    }

    return cfg, nil
}