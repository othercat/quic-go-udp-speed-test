# quic-go-udp-speed-test
use `quic-go` to test udp transport speed, use <https://github.com/quic-go/quic-go>

### Histroy

#### 20240727 version 0.1.0

First version to make client and server work

### Requirement

Client is tested under macOS Sonoma 14.5 with M1 Macbook Air, and Server is tested under CentOS 7.9 x86_64

### Build+Install

1. Just fetch the code via git, then build it with Go

2. Make sure that Server.crt and Server.key are generated for TLS. You can take files from Folder "certs" for reference. 

### Running

Just run "quic-client" or "quic-server" for testing