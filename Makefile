default: build

build:
	@go build -o ./bin/slowloris *.go

run: build
	@./bin/slowloris

server:
	@bash server.sh & > /dev/null 2>&1

kill:
	@echo $(cat server.pid)


