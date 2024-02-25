package main

import (
	"flag"
	"log"
	"os"

	"smtphoneypot/session"
	"smtphoneypot/srvr"
)

func main() {
	hn, _ := os.Hostname()
	hostName := flag.String("hn", hn, "host name to present to senders")
	address := flag.String("a", "localhost:8976", "hostname:port form address on which to listen")
	logDirectory := flag.String("l", "./log", "log directory")
	debug := flag.Bool("d", false, "debug output")

	flag.Parse()

	server, err := srvr.NewServer(*address, *hostName, *logDirectory, *debug)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := server.NextConnection()
		if err != nil {
			server.Debugf("making a connection: %v\n", err)
			continue
		}

		sess, err := session.New(conn, *debug, server.Debugf)
		go sess.Receive()
	}
}
