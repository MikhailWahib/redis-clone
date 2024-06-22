build:
	@go build -o bin/run *.go

run: build
	@./bin/run
