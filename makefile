PB=gen/toupper.pb.go

all: client server

$(PB):
    protoc -I proto toupper.proto --go_out=plugins=grpc:.
server: $(PB)
	go build main/server.go
client: $(PB)
	go build main/client.go

clean:
	rm -f client server

