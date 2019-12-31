#!/usr/bin/env bash

protoc  -I=/Users/qinshen/go/src -I=/usr/local/include  -I=./ --gofast_out=plugins=grpc:.  ./*.proto

flatc --proto  ./protoums.proto

flatc  --go --gen-object-api --gen-all  --gen-mutable --grpc  --gen-compare  --raw-binary ./*.fbs
flatc  --go --gen-object-api --gen-all  --gen-mutable --grpc  --gen-compare  --raw-binary --gen-object-api --size-prefixed --force-empty ./*.fbs


