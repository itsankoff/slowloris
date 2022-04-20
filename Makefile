default: build

build:
	@go build -o ./bin/slowloris *.go

run: build
	@./bin/slowloris
