package libs

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jpillora/overseer"
	"github.com/jpillora/overseer/fetcher"
)

// Fungsi yang dijalankan oleh overseer
// prog(state) runs in a child process
func prog(state overseer.State) {
	log.Printf("app (%s) listening...", state.ID)
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "app (%s) says hello\n", state.ID)
	}))
	http.Serve(state.Listener, nil)
}

// TODO Fetch binaries file
// ? Current Repository
// go build -o tmp/main main.go

// ? Thirdparty
/*
	http.Handle("/", http.FileServer(http.Dir("./tmp")))
	fmt.Println("Listening on port 4000")

	http.ListenAndServe(":4000", nil)
*/
// go build -o tmp/main main.go
// Download binary
// wget -O my_app http://localhost:[PORT]/tmp/main
func Overseer() {
	overseer.Run(overseer.Config{
		Program: prog,
		Address: ":3000",
		Fetcher: &fetcher.File{Path: "./tmp/main"},
	})
}
