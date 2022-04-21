default: build

build:
	@go build -o ./bin/slowloris *.go

run: build
	@./bin/slowloris

server:
	@bash server.sh

kill:
	@cat .server.pid | xargs kill
