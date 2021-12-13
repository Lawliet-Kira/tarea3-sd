gen:
	protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative  helloworld/helloworld.proto

clean:
	rm helloworld/*.go

broker:
	go run Broker/broker.go 
	
server:
	go run lider/lider.go

gamenode:
	go run namenode/gamenode.go
	
datanode:
	go run namenode/datanode.go