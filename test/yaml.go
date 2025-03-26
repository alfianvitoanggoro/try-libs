package test

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func Yaml() {
	// Baca file YAML
	data, err := os.ReadFile(".config.yaml")
	if err != nil {
		log.Fatalf("Gagal membaca file: %v", err)
	}

	// Parse YAML
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Gagal decode YAML: %v", err)
	}

	// Tampilkan hasil
	fmt.Printf("App Name: %s\n", config.App.Name)
	fmt.Printf("Version: %s\n", config.App.Version)
	fmt.Printf("Debug Mode: %t\n", config.App.Debug)
}
