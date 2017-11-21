package main

import (
	"config"
	"pjrouter"
)

func main() {
	router := pjrouter.Load()
	router.RunTLS(":8081", config.CertFile, config.KeyFile)
}
