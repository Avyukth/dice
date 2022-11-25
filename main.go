package main

import (
	"flag"
	"log"

	"dice/config"
	"dice/server"
)

func setupFlags(){
	flag.StringVar(&config.Host, "host", "0.0.0.0", "Host for dice Server")
	flag.IntVar(&config.Port, "port", 7379, "port for dice server")
	flag.Parse()
}

func main(){
	setupFlags()
	log.Println("rolling dice ...")
	server.RunSyncTCPServer()
}
