package httpd

import (
	"net"
	"log"
	"net/http"
)

type Service struct {
	addr     string
	listener net.Listener
}

// Return a new HTTP service.
func New(addr string) *Service {
	return &Service{
		addr: addr,
	}
}

// Start the HTTP service.
func (s *Service) Start() error {
	server := http.Server{
		Handler: s,
	}

	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	s.listener = listener

	http.Handle("/", s)

	go func() {
		err := server.Serve(s.listener)
		if err != nil {
			log.Fatalf("HTTP serve: %s", err)
		}
	}()

	log.Printf("Started HTTP service at %s", s.addr)
	return nil
}

// Close the service
func (s *Service) Close() {
	s.listener.Close()
}

// Serve the HTTP request (implements the Handler interface).
func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got %s request at %s",
		r.Method,
		r.URL.Path)
}

// Addr returns the address on which the Service is listening
func (s *Service) Addr() net.Addr {
	return s.listener.Addr()
}
