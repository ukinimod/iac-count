package main

import (
	"log"

	"github.com/MaibornWolff/iac-count/cmd"
)

func main() {
	err := cmd.RootCmd.Execute()

	if err != nil {
		log.Fatal(err)
	}
}
