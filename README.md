# Simple_Distributed_System (Doc unfinished)
This project is a basic distributed system used to record the data of baseball game.
Using gRPC (Golang), replica-set (mongoDB), Cobra (CLI) this project could ...
1. Trigger the function with gRPC.
2. Create or update the the baseball game record with CLI terminal.
3. Backup database automatically.

## Quick Start Server
![](https://raw.githubusercontent.com/amosricky/Simple_Distributed_System/master/src/server_node.gif)
```
$ cd $GOPATH/src

# Git clone from github
$ git clone https://github.com/amosricky/Simple_Distributed_System.git

# Change path to project
$ cd ./Simple_Distributed_System

# Run it 
$ docker-compose -f server_node.yml up
```

## Quick Start Client
```
# Change path to project
$ cd ./Simple_Distributed_System

# Run it 
$ docker-compose -f client_node.yml run client_node 
```

## Features


## How it works
```
.
├── client
│   ├── clientNode.go
│   └── cmd
│       └── cmd.go
├── client_node
├── client_node.yml
├── conf
│   └── app.ini
├── go.mod
├── go.sum
├── LICENSE
├── pb
│   ├── game.pb.go
│   └── game.proto
├── README.md
├── replica
│   ├── config
│   │   └── mongod.yml
│   └── data
│       ├── rs1
│       ├── rs2
│       └── rs3
├── replica_set.js
├── server
│   └── serverNode.go
├── server_node
├── server_node.sh
├── server_node.yml
└── setting
    └── setting.go
...
```

### Required

- Golang >= 1.13.7
- docker-compose
- mongo >= 4.2




protoc --go_out=plugins=grpc:. *.proto
docker-compose -f client_node.yml run client_node
https://blog.ruanbekker.com/blog/2019/04/17/mongodb-examples-with-golang/
