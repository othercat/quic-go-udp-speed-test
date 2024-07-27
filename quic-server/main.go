package main

import (
    "crypto/tls"
    "log"
    "net"
    "net/http"
    "github.com/quic-go/quic-go/http3"
)

func main() {
    // create UDP listener
    lis, err := net.ListenPacket("udp", ":4242")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    // load certs and keys
    tlsConfig := &tls.Config{
        Certificates: []tls.Certificate{},
    }
    cert, err := tls.LoadX509KeyPair("server_127.0.0.1.crt", "server_127.0.0.1.key")
    if err != nil {
        log.Fatalf("failed to load key pair: %v", err)
    }
    tlsConfig.Certificates = append(tlsConfig.Certificates, cert)

    // create simeple HTTP server
    handler := http.FileServer(http.Dir("."))

    // create HTTP/3 server
    http3Server := &http3.Server{
        Handler:    handler,
        TLSConfig:  tlsConfig,
    }

    // launch HTTP/3 server
    err = http3Server.Serve(lis)
    if err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
