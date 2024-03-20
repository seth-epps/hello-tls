package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

type serverConfig struct {
	tlsKeyPath  string
	tlsCertPath string
	port        int
}

type Response struct {
	IP      string `json:"ip"`
	Message string `json:"message"`
}

func main() {
	var server serverConfig
	flag.StringVar(&server.tlsKeyPath, "tls-key", "/etc/certs/tls.key", "Path to the TLS key")
	flag.StringVar(&server.tlsCertPath, "tls-cert", "/etc/certs/tls.crt", "Path to the TLS certificate")
	flag.IntVar(&server.port, "port", 8443, "Server port")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handleGet)

	addr := fmt.Sprintf(":%d", server.port)

	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	log.Printf("Server starting on %s", addr)
	log.Fatal(srv.ListenAndServeTLS(server.tlsCertPath, server.tlsKeyPath))
}

func handleGet(w http.ResponseWriter, req *http.Request) {
	res := Response{IP: req.RemoteAddr, Message: "Hello From Go w/ TLS!"}
	json.NewEncoder(w).Encode(res)
}
