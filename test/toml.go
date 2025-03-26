package test

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

func Toml() {
	// Read file .config.toml
	data, err := os.ReadFile(".config.toml")
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	fmt.Printf("Data: %s\n", data)

	// Parse TOML
	var config *Config
	err = toml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error parsing TOML:", err)
		return
	}

	// Cetak versi
	fmt.Println("Version:", config.App.Name)
	fmt.Println("Version:", config.App.Version)
	fmt.Println("Version:", config.App.Debug)
}
