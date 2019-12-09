package main

import (
	"os"
	"serve/api"
	flag "github.com/spf13/pflag"
)

const (
	PORT string = "8000"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}

	pPort := flag.StringP("port", "p", PORT, "PORT for http to listen")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	server := api.NewServer()
	server.Run(":" + port)
}
