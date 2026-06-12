package server

import (
	"bufio"
	"net"
	"strings"
)

func (s *TCPServer) handleConnection(
	conn net.Conn,
) {

	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {

		input := strings.TrimSpace(
			scanner.Text(),
		)

		response := s.Execute(input)

		conn.Write(
			[]byte(response + "\n"),
		)
	}
}

func (s *TCPServer) Execute(
	input string,
) string {

	cmd := ParseCommand(input)

	switch cmd.Action {

	case "SET":

		s.store.Set(
			cmd.Key,
			cmd.Value,
		)

		return "OK"

	case "GET":

		value, found :=
			s.store.Get(cmd.Key)

		if !found {
			return "NOT FOUND"
		}

		return value

	case "DEL":

		s.store.Delete(cmd.Key)

		return "DELETED"

	case "EXISTS":

		if s.store.Exists(cmd.Key) {
			return "TRUE"
		}

		return "FALSE"

	default:

		return "INVALID COMMAND"
	}
}