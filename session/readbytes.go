package session

func (s *Session) readCommand() []byte {

	buffer := make([]byte, 1024)

	foundEndOfLine := false
	endCommand := -1

	for !foundEndOfLine {
		n, err := s.conn.Read(buffer)
		s.debugf("read %d bytes into buffer\n", n)
		if err != nil {
			s.debugf("readCommand: %v\n", err)
			break
		}
		if n == 0 {
			break
		}

		s.buffer = append(s.buffer, buffer[:n]...)

		foundCR := false
		for i := range s.buffer {
			if s.buffer[i] == '\r' {
				s.debugf("found CR at offset %d\n", i)
				foundCR = true
				continue
			}
			if foundCR && s.buffer[i] == '\n' {
				s.debugf("found LF at offset %d\n", i)
				endCommand = i
				foundEndOfLine = true
			}
		}
	}
	if endCommand > 0 {
		s.debugf("end-of-command at offset %d\n", endCommand)
		command := s.buffer[0 : endCommand-1] // trim off CRLF
		s.buffer = s.buffer[endCommand+1:]
		s.debugf("command is %d bytes, buffer contains %d leftover bytes\n", len(command), len(s.buffer))
		return command
	}
	return nil
}
