package main

import (
	"fmt"

	"github.com/mbStavola/tarkov-api"
)

func main() {
	api := tarkov.NewTarkovAPI("dummy-session-id")

	bosses, err := api.Characters().Bosses()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(bosses)
}
