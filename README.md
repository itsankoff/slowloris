# slowloris - Golang distributed Slowloris attack

![Cute Slowloris](slowloris.png)


## How it works
Read more [here](https://thetooth.io/blog/slowloris-attack) 🦷

## How to protect from it
TBD

## Installation
* Make sure you have golang installation `1.16+`
* Run `go install github.com/itsankoff/slowloris`

## Usage

**WARNING**:
This software MUST NOT BE used for malicious purpose that may cause harm on
any third party. Use it only for educational purposes and at own discretion.

* Basic usage: `slowloris -url=<URL>`
* For more sophisticated usage use `slowloris help` to get the full option set

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
* [Wikipedia](https://en.wikipedia.org/wiki/Slowloris_(computer_security))
* [CloudFlare](https://www.cloudflare.com/learning/ddos/ddos-attack-tools/slowloris/)

## License
* [MIT](LICENSE)
