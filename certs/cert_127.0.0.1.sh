openssl genrsa -out server_127.0.0.1.key 2048
openssl req -new -key server_127.0.0.1.key -out server_127.0.0.1.csr -config openssl_127.0.0.1.cnf
openssl x509 -req -days 365 -in server_127.0.0.1.csr -signkey server_127.0.0.1.key -out server_127.0.0.1.crt -extensions v3_ca -extfile openssl_127.0.0.1.cnf
