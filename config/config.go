package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)


type Config struct {
	Port string `mapstructure:"PORT"`
}

func LoadEnvVars() (Config, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileAlreadyExistsError); !ok {
			return Config{}, fmt.Errorf("error loading config file: %s", err.Error())
		}
	}

	_ = viper.BindEnv("PORT")

	for _, key := range viper.AllKeys() {
		val := viper.Get(key)
		viper.Set(key, val)
	}

	c := Config{}
	err = viper.Unmarshal(&c)
	if err != nil {
		return Config{}, fmt.Errorf("error while unmarshalling config file: %w", err)
	}

	v := validator.New()
	err = v.Struct(c)
	if err != nil {
		return Config{}, fmt.Errorf("invalid config file: %w", err)
	}
	
	return c, nil
}