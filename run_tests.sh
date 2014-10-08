# Kill any existing MongoDB, otherwise we'd accidentally use
# an unintended instance for test
pkill mongod

# Startup Test MongoDB Instance
nohup mongod --dbpath=./tests/mongo &
go fmt ./...
go test ./...
pkill mongod