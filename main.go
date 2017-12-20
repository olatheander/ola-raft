package main

import (
	"flag"
	"os"
	"os/signal"
	"log"
	"github.com/olatheander/ola-raft/httpd"
)

// Command line defaults
const (
	DefaultHTTPAddr = ":11000"
)

// Command line parameters
var nodeId string
var httpAddr string

func init() {
	flag.StringVar(&httpAddr, "haddr", DefaultHTTPAddr, "Set the HTTP bind address")
	flag.StringVar(&nodeId, "id", "<missing>", "Node ID")
}

func main() {
	flag.Parse()

	h := httpd.New(httpAddr)
	if err := h.Start(); err != nil {
		log.Fatalf("failed to start the HTTP service: %s", err.Error())
	}

	log.Printf("ola-raft started successfully with nodeId: %s\n", nodeId)

	terminate := make(chan os.Signal, 1)
	signal.Notify(terminate, os.Interrupt)
	<-terminate
	log.Printf("ola-raft exiting")
}
