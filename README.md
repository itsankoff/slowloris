# slowloris - Golang distributed Slowloris attack

![Cute Slowloris](assets/slowloris.png)


## How it works
Read the [article](https://thetooth.io/blog/slowloris-attack) 🦷

## How to protect from it
TBD

## Installation
* Run `go install github.com/itsankoff/slowloris`

## Usage
* Basic usage
```
slowloris -url="http[s]://<domain>[:<port>]/<path>?<query-string>"
slowloris -url="https://example.com"
```

* For more sophisticated usage use to get the full option set:
```
slowloris help
```

* **LEGAL DISCLAIMER**
```
Usage of this program for attacking targets without
prior mutual consent is illegal. It is the end user's responsibility to obey
all applicable local, state and federal laws in all countries.
Developers assume no liability and are not responsible for any misuse or
damage caused by this program.
```

## Testing

The Makefile support simple HTTP server that you can use for testing purposes.
* Start the server:
```
# The server listens on localhost:8080 and runs in background
# Log of the server in ./.server.log
# PID of the server in ./.server.pid
make server
```
* Stop the server:
```
make kill
```
* Get statistics about `ESTABLISHED` connections:
```
make stats
```
* Get statistics about count of connections in other states:
```
make stats STATE=<STATE (e.g. LISTEN)>
```

## Reference
* [thetooth.io](https://thetooth.io/blog/slowloris-attack/)
* [Wikipedia](https://en.wikipedia.org/wiki/Slowloris_(computer_security))
* [CloudFlare](https://www.cloudflare.com/learning/ddos/ddos-attack-tools/slowloris/)

## License
[MIT](LICENSE)
