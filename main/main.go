package main

import (
	"goworkshop/web"
	"goworkshop/persistence"
)

func main() {
	if err := persistence.InitDB(); err != nil {
		panic(err)
	}

	server := web.RestServer{
		Port: 8080,
	}

	server.StartServer()
}
