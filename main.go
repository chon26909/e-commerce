package main

import (
	"os"

	"github.com/chon26909/e-commerce/config"
	"github.com/chon26909/e-commerce/modules/server"
	"github.com/chon26909/e-commerce/pkg/database"
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

	db := database.NewDatabase(config.Db())
	defer db.Close()

	server.NewServer(config, db).Start()

}
