package main

import (
	"github.com/cmelgarejo/go-gql-server/Config"
	"github.com/cmelgarejo/go-gql-server/server"
)

func main() {
	server.Run()
	Config.start()
	Config.UserModel
}
