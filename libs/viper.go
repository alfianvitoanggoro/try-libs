package libs

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var v *viper.Viper

func Viper() {
	// Read file
	v = viper.New()
	v.SetConfigName(".config") // Name of config file
	v.SetConfigType("toml")    // Extension of file
	v.AddConfigPath(".")       // optionally look for config in the working directory

	if err := v.ReadInConfig(); err != nil {
		log.Println(err.Error())
		return
	}

	// Val config variable to hold the config value.
	var Val *Config

	if err := v.Unmarshal(&Val); err != nil {
		log.Println(err.Error())
	}

	fmt.Printf("App Name: %s\n", Val.App.Name)
	fmt.Printf("Version: %s\n", Val.App.Version)
	fmt.Printf("Debug Mode: %t\n", Val.App.Debug)
}
