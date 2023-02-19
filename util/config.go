package util

import (
	"fmt"

	"github.com/spf13/viper"
)

func NewConfig(path string) (*viper.Viper, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error when read config: %v", err)
	}

	return viper.GetViper(), nil
}
