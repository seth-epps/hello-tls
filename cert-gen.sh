#!/bin/bash

######
####
## This is for local testing only!
####
######

CERTS_DIR=".certs"

if [[ $# == 0 ]]; then
  echo "Create self signed CA and server domain certificates in $CERTS_DIR"
  echo "Usage: $0 SAN [SAN...]" >&2
  echo "" >&2
  echo "Currently only supports DNS SANs"
  echo "" >&2
  exit 1
fi

COMMON_NAME=$1
FORMATTED_SANS=$(printf 'DNS:%s,' "${@:1}" | sed 's/,$/\n/')

# create $CERTS_DIR
mkdir -p $CERTS_DIR

# # Create private key for server and self-sign
openssl genrsa -out $CERTS_DIR/server.key 4096
openssl req -new -key $CERTS_DIR/server.key -sha256  -out $CERTS_DIR/server.csr -subj "/CN=${COMMON_NAME}"
openssl x509 -req -extfile <(printf "subjectAltName = $FORMATTED_SANS") -in $CERTS_DIR/server.csr -signkey $CERTS_DIR/server.key -out $CERTS_DIR/server.crt
