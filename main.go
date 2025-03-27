package main

import (
	"fmt"

	"github.com/alfianvitoanggoro/try-libs/cmd"
	"github.com/alfianvitoanggoro/try-libs/test"
)

func main() {
	cmd.Execute()
	test.Test()
	fmt.Println("test")
}
