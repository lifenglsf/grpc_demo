gen:
	protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=. 
clean:
	rm pb/*.go
run:
	go run main.go
server:
	go run cmd/server/main.go -port 9998
client:
	go run cmd/client/main.go -address 0.0.0.0:9998
test:
	go test --cover --race ./...