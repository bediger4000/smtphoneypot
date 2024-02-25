package session

import (
	"net"
)

type Session struct {
	conn   net.Conn
	debug  bool
	debugf func(string, ...any)
}

func New(conn net.Conn, debug bool, debugf func(string, ...any)) (*Session, error) {
	return &Session{
		conn:   conn,
		debug:  debug,
		debugf: debugf,
	}, nil
}

func (s *Session) Receive() {
	s.debugf("closing connection to %s\n", s.conn.RemoteAddr())
	s.conn.Close()
}
