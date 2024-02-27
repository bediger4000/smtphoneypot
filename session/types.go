package session

import (
	"net"
)

type Session struct {
	conn   net.Conn
	debug  bool
	debugf func(string, ...any)
	buffer []byte
}

func New(conn net.Conn, debug bool, debugf func(string, ...any)) (*Session, error) {
	return &Session{
		conn:   conn,
		debug:  debug,
		debugf: debugf,
	}, nil
}

func (s *Session) Receive() {
	s.debugf("receiving from %s\n", s.conn.RemoteAddr())
	for {
		line := s.readCommand()
		s.debugf("received %v\n", line)
		if len(line) == 0 {
			break
		}
	}
	s.debugf("closing connection to %s\n", s.conn.RemoteAddr())
	s.conn.Close()
}
