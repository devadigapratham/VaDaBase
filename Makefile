build: 
	@go build -o bin/VaDaBase cmd/main.go 

run: build
	@./bin/VaDaBase

test: 
	@go test -v ./... 
	