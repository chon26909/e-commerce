package main

import (
	"fmt"
	"os"

	"github.com/chon26909/e-commerce/config"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

func main() {
	config := config.LoadConfig(envPath())

	fmt.Println(config)
}
