PROTO_DIR = proto-go
PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')

build: generate
	go build cmd/main.go

generate:
	protoc -I${PROTO_DIR} --go_opt=module=${PACKAGE}/proto-go --go_out=./proto-go ${PROTO_DIR}/*.proto

bump: generate
	go get -u ./...

clean:
	rm ${PROTO_DIR}/*.pb.go
	rm main
