package srvr

import (
	"fmt"
	"net"
	"os"
	"time"
)

type SMTPServer struct {
	address      string
	lstnr        net.Listener
	hostName     string
	logDirectory string
	debug        bool
	serverLog    *os.File
}

func NewServer(listenAddress, hostName, logDirectory string, debug bool) (*SMTPServer, error) {

	lstnr, err := net.Listen("tcp", listenAddress)
	if err != nil {
		return nil, err
	}

	logout, err := createServerLog(logDirectory)
	if err != nil {
		return nil, err
	}

	return &SMTPServer{
		address:      listenAddress,
		lstnr:        lstnr,
		hostName:     hostName,
		logDirectory: logDirectory,
		debug:        debug,
		serverLog:    logout,
	}, nil
}

func (s *SMTPServer) NextConnection() (net.Conn, error) {
	s.Debugf("accepting new connection on %s\n", s.address)
	conn, err := s.lstnr.Accept()
	if err != nil {
		return nil, err
	}
	if s.debug {
		s.Debugf("connection from %s\n", conn.RemoteAddr())
	}
	return conn, err
}

func createServerLog(logDirectory string) (*os.File, error) {
	return os.Stderr, nil
}

func (s *SMTPServer) Debugf(formatString string, a ...any) {
	if !s.debug {
		return
	}
	fmt.Fprintf(
		s.serverLog,
		fmt.Sprintf("%s\t%s", time.Now().Format(time.RFC3339), formatString),
		a...,
	)
}
