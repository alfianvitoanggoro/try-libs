package libs

import (
	"bytes"
	"fmt"
	"os"

	"github.com/vmihailenco/msgpack/v5"
)

var (
	MsgpackStruct    bool
	MsgpackMap       bool
	MsgpackStreaming bool
	MsgpackJSON      bool
)

type Msgpack struct {
}

func NewMsgpack() *Msgpack {
	return &Msgpack{}
}

func (m *Msgpack) Msgpack() {
	// Data yang akan dikodekan
	user := User{Name: "John Doe", Age: 30, Email: "john@example.com"}

	// Encode ke format MessagePack
	data, err := msgpack.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println("Encoded Data:", data)

	// Decode kembali ke struct
	var decodedUser User
	err = msgpack.Unmarshal(data, &decodedUser)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decoded Struct:", decodedUser)
}

func (m *Msgpack) MsgpackMap() {
	// Data berbentuk map
	dataMap := map[string]interface{}{
		"id":    1,
		"name":  "Alice",
		"score": 95.5,
	}

	// Encode ke MessagePack
	encoded, _ := msgpack.Marshal(dataMap)
	fmt.Println("Encoded Map:", encoded)

	// Decode kembali
	var decodedMap map[string]interface{}
	msgpack.Unmarshal(encoded, &decodedMap)

	fmt.Println("Decoded Map:", decodedMap)
}

func (m *Msgpack) MsgpackStreaming() {
	// Simpan data ke file
	file, _ := os.Create("data.msgpack")
	defer file.Close()

	encoder := msgpack.NewEncoder(file)

	// Encode data ke file
	user := User{Name: "Alice", Age: 25, Email: "alice@example.com"}
	encoder.Encode(user)

	//  * Decode
	// Open file untuk decoding
	file, err := os.Open("data.msgpack")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Pastikan pointer di awal file
	file.Seek(0, 0)

	// Decode dari file
	decoder := msgpack.NewDecoder(file)
	var decodedUser User
	err = decoder.Decode(&decodedUser)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}

	fmt.Println("Decoded User:", decodedUser)
}

func (m *Msgpack) MsgpackJSON() {
	user := User{Name: "John Doe", Age: 30, Email: "john@example.com"}

	var buf bytes.Buffer
	enc := msgpack.NewEncoder(&buf)
	// ? Using custom tag for json
	enc.SetCustomStructTag("json")

	enc.Encode(user)
	fmt.Println("Encoded Data:", buf.String())
}
