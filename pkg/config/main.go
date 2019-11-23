package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func init() {
	env := os.Getenv("APP_ENV")
	conf := "config"
	if env != "" {
		conf += "." + env
	}
	viper.SetConfigName(conf)     // name of config file (without extension)
	viper.SetConfigType("yaml")   // or viper.SetConfigType("YAML")
	viper.AddConfigPath("./conf") // optionally look for config in the working directory
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

}

func GetEnv() string {
	v := os.Getenv("APP_ENV")
	return v
}

// MustGet will return the env or panic if it is not present
func MustGet(k string) string {
	var v string
	if viper.IsSet(k) {
		v = viper.GetString(k)
	} else {
		log.Panicln("config missing, key: " + k)
	}
	return v
}

// MustGetBool will return the env as boolean or panic if it is not present
func MustGetBool(k string) bool {
	var v bool
	if viper.IsSet(k) {
		v = viper.GetBool(k)
	} else {
		log.Panicln("config missing, key: " + k)
	}
	return v
}
