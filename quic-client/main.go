package main

import (
    "github.com/quic-go/quic-go/http3"
    "io"
    "log"
    "net/http"
    "os"
    "time"
)

func main() {
    // create HTTP/3 client config
    roundTripper := &http3.RoundTripper{}

    // create HTTP client
    client := &http.Client{
        Transport: roundTripper,
    }

    // initial start time
    start := time.Now()

    // send download request
    resp, err := client.Get("https://127.0.0.1:4242/small-file-1MB.bin") //  large-file-100MB.bin
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    // create local file
    outFile, err := os.Create("Get-file.bin")
    if err != nil {
        log.Fatalf("Failed to create file: %v", err)
    }
    defer outFile.Close()

    // copy resp to local file
    _, err = io.Copy(outFile, resp.Body)
    if err != nil {
        log.Fatalf("Failed to save file: %v", err)
    }

    // count the download time
    elapsed := time.Since(start)
    log.Printf("Download took %s", elapsed)
}
