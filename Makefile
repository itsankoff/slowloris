default: build

SERVER_PID=.server.pid

build:
	@go build -o ./bin/slowloris *.go

run: build
	@./bin/slowloris

server:
	@test -f ${SERVER_PID} || bash server.sh

kill:
	@cat .server.pid | xargs kill
	@rm .server.pid
