# Archetype of a Go project that can be built with Nix

This project is still **Work In Progress**. I think I've got most of
the necessary dependencies lined up, but I still have to connect the  
dots...

## Introduction

My ultimate aim is to create a project that
can be used as a template for future projects with the following
characteristics:
* Command / Query Responsibility Segregation (CQRS)
* Event Sourcing
* Strong typing
* High Performance
* High Availability
* Scalability

On top of that, I would love to be able to start a project as a monolithic
application and have it evolve into a collection of micro-services that is
integrated into a service mesh architecture.

The components that I want to combine are:
* The Go language (for high-performance, type safety and gRPC integration with Axon Server)
* Docker (to minimise the impact on/from the host system)
* Nix (to manage dependencies)
* Axon Server (for event storage, message routing, and scalability)
* Envoy (for service mesh architecture and high availability)

This project is a template for a gRPC enabled Go back-end that can be built with Nix.
The saga continues in [archetype-go-axon](https://github.com/jeroenvanmaanen/archetype-go-axon)

## Setup

I mostly followed [Golang Demo](https://github.com/MatrixAI/Golang-Demo)
by _Roger Qiu_. Any flaws are of course my own.

To work with this project, you need to install docker. The first step after
that is to acquire a docker image that has Nix and Go tools. It will be pulled from
docker hub automatically the first time you run `docker run` or
`docker-compose up`. You can also build it yourself with:
```
[host]$ src/bin/docker-build.sh # Optional. It is also available on Docker Hub
```

You might need to update `deps.nix`:
```
[host]$ docker run --rm -ti -v "${HOME}:${HOME}" -w "$(pwd)" jeroenvm/archetype-nix-go bash
[container]# vgo2nix
```

After that, build the executables from the Go code:
```
[host]$ src/bin/nix-build.sh
```

I generated Go stubs for axon-server as follows:
```
[host]$ docker run --rm -ti -v "${HOME}:${HOME}" -w "$(pwd)" jeroenvm/archetype-nix-go bash
[container]# go get -u github.com/golang/protobuf/protoc-gen-go
[container]# go get google.golang.org/grpc
[container]# PATH="$PATH:/root/go/bin"
[container]# cd /src/axon-server-api/src/main/proto
[container]# bash WORKING_AREA/archetype-nix-go/src/bin/generate-proto-package.sh
```

To run the whole setup, execute the bash script `src/docker/docker-compose-up.sh`.
Then from the root of the project in another terminal window:
```
[host]$ docker exec -ti -w "$(pwd)" example_example_1 bash
[container]# result/bin/example
```
The last command outputs something like:
```
3
primary:<host_name:"7b43732c1846" grpc_port:8124 http_port:8024 node_name:"7b43732c1846" > same_connection:true 
xxx
```
which proves that it could reach Axon Server through gRPC.
