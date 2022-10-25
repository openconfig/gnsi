#!/bin/bash
set -euo pipefail

proto_imports=".:${GOPATH}/src/github.com/google/protobuf/src:${GOPATH}/src/github.com/googleapis/googleapis:${GOPATH}/src:."

# Go
for p in accounting/acct.proto authz/authz.proto certz/certz.proto credentialz/credentialz.proto pathz/pathz.proto pathz/authorization.proto; do
  protoc -I="${proto_imports}" --go-grpc_out=. --go-grpc_opt=paths=source_relative --go_out=. --go_opt=paths=source_relative $p
done

