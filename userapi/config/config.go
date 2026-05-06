package config

import "github.com/spf13/viper"

type Config struct {
    Server   ServerConfig
    Database DatabaseConfig
    JWT      JWTConfig
}

type ServerConfig struct {
    Port int
}

type DatabaseConfig struct {
    Path string
}

type JWTConfig struct {
    Secret string
    Expire int
}

func Load() (*Config, error) {
    viper.SetConfigName("config")  // 配置文件名
    viper.SetConfigType("yaml")    // 配置文件类型
    viper.AddConfigPath(".")       // 查找路径

    if err := viper.ReadInConfig(); err != nil {
        return nil, err
    }

    var cfg Config
    if err := viper.Unmarshal(&cfg); err != nil {
        return nil, err
    }
    return &cfg, nil
}