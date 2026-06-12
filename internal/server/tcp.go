package server


import (
	"fmt"
	"net"
	"github.com/nishanth-code/redis-clone/internal/store"
)

type TCPServer struct {
	port  string
	store *store.Store
}

func NewTCPServer(
	port string,
	store *store.Store,
) *TCPServer {
	return &TCPServer{
		port:  port,
		store: store,
	}
}

func (s *TCPServer) Start() error {

	listener, err := net.Listen(
		"tcp",
		":"+s.port,
	)

	if err != nil {
		return err
	}

	defer listener.Close()

	fmt.Println("TCP Server listening on port", s.port)

	for {

		conn, err := listener.Accept()

		if err != nil {
			fmt.Println(err)
			continue
		}

		go s.handleConnection(conn)
	}
}