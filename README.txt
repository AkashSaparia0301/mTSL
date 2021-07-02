Install Openssl in desktop (https://slproweb.com/products/Win32OpenSSL.html)

# CA key and certificate
openssl genrsa -des3 -out ca.key 4096
openssl req -new -x509 -days 365 -key ca.key -out ca.crt

# server key
openssl genrsa -des3 -out server.key 1024

# CSR (certificate sign request) to obtain certificate
openssl req -new -key server.key -out server.csr

# sign server CSR with CA certificate and key
openssl x509 -req -days 365 -in server.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out server.crt

# client key
openssl genrsa -des3 -out client.key 1024

# CSR to obtain certificate
openssl req -new -key client.key -out client.csr

# sign client CSR with CA certificate and key
openssl x509 -req -days 365 -in client.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out client.crt

# server key out to temp.key
openssl rsa -in server.key -out temp.key

# remove server.key
rm server.key (Use "del" in CMD instead of "rm")

# make temp.key as server key
mv temp.key server.key (Use "move" in CMD instead of "mv")

# client key out to temp.key
openssl rsa -in client.key -out temp.key

# remove client.key
rm client.key (Use "del" in CMD instead of "rm")

# make temp.key as client key
mv temp.key client.key (Use "move" in CMD instead of "mv")
