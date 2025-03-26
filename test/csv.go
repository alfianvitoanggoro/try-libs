package test

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

type CSV struct {
}

func NewCSV() *CSV {
	return &CSV{}
}

func (c *CSV) ReadCSV() {
	clientsFile, err := os.OpenFile(".config.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer clientsFile.Close()

	clients := []*Client{}

	if err := gocsv.UnmarshalFile(clientsFile, &clients); err != nil { // Load clients from file
		panic(err)
	}

	for _, client := range clients {
		fmt.Printf("%d. My Name is %s and I am %d years old\n", client.ID, client.Name, client.Age)
	}

	if _, err := clientsFile.Seek(0, 0); err != nil { // Go to the start of the file
		panic(err)
	}
}

func (c *CSV) CreateDataCSV() {
	// Data yang akan disimpan
	clients := []*Client{
		{1, "Alice", 50},
		{2, "Bob", 42},
	}

	// Buat file CSV
	file, err := os.Create(".config.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Tulis struct ke CSV
	if err := gocsv.MarshalFile(&clients, file); err != nil {
		panic(err)
	}

	fmt.Println("Data berhasil disimpan ke .config.csv")
}
