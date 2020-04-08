gen:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:framework/pb

clean:
	rm pb/*.go

server:
	go run framework/cmd/server/main.go -port 8080

client:
	go run framework/cmd/client/main.go -address 0.0.0.0:8080

test:
	go test -cover -race ./...

.PHONY: gen clean server client test