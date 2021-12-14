gen:
	protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative  helloworld/helloworld.proto

clean:
	rm planetas/*.txt
	rm logs/*.txt

broker:
	go run Broker/broker.go 
	
info:
	go run Informantes/informante.go

fulcrum:
	go run Fulcrums/fulcrum.go

git:
	git add .
	git commit -m "a"
	git push