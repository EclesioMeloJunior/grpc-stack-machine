PRJ_DIR=./
DST_DIR=$(PRJ_DIR)

.PHONY: proto
proto:
	protoc -I=$(PRJ_DIR) --go_out=plugins=grpc:$(DST_DIR) $(PRJ_DIR)/machine/machine.proto

run:
	go run main.go -port 9000

test:
	go test ./...