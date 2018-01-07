package main

import (
	"flag"
	"helloserver/server"
	"helloserver/utils"
	"os"
)

var (
	ip       = flag.String("ip", "http://localhost", "ip address to listen")
	port     = flag.String("port", "8080", "port to listen")
	greeting = flag.String("greeting", "", "Greeting message to response")
	logger   = utils.NewLogger()
)

func main() {
	flag.Parse()
	// args =

	if *greeting == "" {
		logger.Println("Usage: helloserver --greeting <message>")
		os.Exit(1)
	}

	config := &utils.Config{Ip: *ip, Port: *port, Greeting: *greeting}

	server := server.NewServer(config)
	server.RunServer()
}
