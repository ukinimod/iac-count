package main

import (
	"log"

	"github.com/ukinimod/iac-count/cmd"
)

func main() {
	err := cmd.RootCmd.Execute()

	if err != nil {
		log.Fatal(err)
	}
}
