#!/bin/sh

mongo library_test --eval "printjson(db.dropDatabase())"
export GIN_MODE=release
go build cmd/server/server.go
./server library_test &
PID=$!
go test ./...
kill $PID
wait $PID
