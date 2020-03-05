#!/usr/bin/env bash

#openssl genrsa -out server.key 2048

# openssl req -new -x509 -sha256 -key server.key \
# -out server.crt -days 3650  -extfile v3.ext
#Create the certificate referencing this config file

openssl req -x509 -nodes -days 3650 -newkey rsa:2048 \
 		-keyout server.key -out server.crt -config req.cnf -sha256

