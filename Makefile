.PHONY: worker client

worker: 
	go run ./worker/server.go

stop: 
	go run ./client/main.go
