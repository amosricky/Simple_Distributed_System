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
![](https://raw.githubusercontent.com/amosricky/Simple_Distributed_System/master/src/client_node.gif)
```
# Change path to project
$ cd ./Simple_Distributed_System

# Run it 
$ docker-compose -f client_node.yml run client_node 
```

## CLI Document
```
$ game help
Get & Modify game record

Usage:
   game [command]

Available Commands:
  add         Add point by game ID.
  list        Get game list. (Contain gameName & gameID)
  new         Create a new game.
  score       Get score by game ID.

Flags:
  -h, --help   help for game
```
```
$ game new -h
Create a new game.

Usage:
   game new [flags]

Flags:
  -h, --help          help for new
  -n, --name string   game name
```
```
$ game list -h
Get game list. (Contain gameName & gameID)

Usage:
   game list [flags]

Flags:
  -d, --dbIP string    database ip (default "127.0.0.1")
  -p, --dbPort int32   database port (default 27041)
  -h, --help           help for list
```
```
$ game add -h
Add point by game ID.

Usage:
   game add [flags]

Flags:
  -a, --add int32     add point
  -h, --help          help for add
  -i, --id string     game id
  -r, --round int32   round ([min]1 [max]9)
  -t, --team int32    team ([0]home [1]visitor) (default -1)
```
```
$ game score -h
Get score by game ID.

Usage:
   game score [flags]

Flags:
  -d, --dbIP string    database ip (default "127.0.0.1")
  -p, --dbPort int32   database port (default 27041)
  -h, --help           help for score
  -i, --id string      game id
```
## Demo
* Create a new game called "testGame", and then check game list.

![](https://raw.githubusercontent.com/amosricky/Simple_Distributed_System/master/src/demo_addGame.gif)

* Add 5 point for home team on second round.
* Add 3 point for guest team on second round.

![](https://raw.githubusercontent.com/amosricky/Simple_Distributed_System/master/src/demo_addScore.gif)

* Read the data from different replica node. (Because we use replica set as database node which could backup the data automatically, we could READ the data from every replica set node.)

![](https://raw.githubusercontent.com/amosricky/Simple_Distributed_System/master/src/demo_getInfo.gif)

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
