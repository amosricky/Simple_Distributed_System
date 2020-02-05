# Simple_Distributed_System
This project is a basic distributed system used to record the data of baseball game.
Using gRPC (Golang), replica-set (mongoDB), Cobra (CLI) this project could ...
1. Trigger the function with gRPC.
2. Create or update the the baseball game record with CLI terminal.
3. Backup database automatically.

## Article
1. [[筆記] 實作分散式計分系統(一) : 基礎架構](https://medium.com/@amosricky95/%E7%AD%86%E8%A8%98-%E5%AF%A6%E4%BD%9C%E5%88%86%E6%95%A3%E5%BC%8F%E8%A8%88%E5%88%86%E7%B3%BB%E7%B5%B1-%E4%BA%8C-replica-set-in-container-5759b1b4cd5)
1. [[筆記] 實作分散式計分系統(二) : Replica Set in Container](https://medium.com/@amosricky95/%E7%AD%86%E8%A8%98-%E5%AF%A6%E4%BD%9C%E5%88%86%E6%95%A3%E5%BC%8F%E8%A8%88%E5%88%86%E7%B3%BB%E7%B5%B1-%E4%BA%8C-replica-set-in-container-5759b1b4cd5)
1. [[筆記] 實作分散式計分系統(三) : Unfinished ](https://medium.com/@amosricky95/%E7%AD%86%E8%A8%98-%E5%AF%A6%E4%BD%9C%E5%88%86%E6%95%A3%E5%BC%8F%E8%A8%88%E5%88%86%E7%B3%BB%E7%B5%B1-%E4%BA%8C-replica-set-in-container-5759b1b4cd5)


## System Architecture Diagram
![](https://raw.githubusercontent.com/amosricky/Simple_Distributed_System/master/src/system_architecture_diagram.png)
1. Each node is a container.
2. Replica set are composed of three mongoDB node - Rs1(Primary), Rs2(Seconary) and Rs3(Seconary).
3. Server node is a gRPC server, the exposed port is 50051.
4. Client node could call RPC function through CLI.

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

* Read the data from different replica node.\
 (Because we use replica set as database node which could backup the data automatically, we could READ the data from every replica set node.)

![](https://raw.githubusercontent.com/amosricky/Simple_Distributed_System/master/src/demo_getInfo.gif)

## Features
1. Use replica set(mongoDB) to backup data automatically.
2. Use protobuf to define the baseball game structure.
3. Use gRPC instead of Restful API which work on http2.
4. Use Cobra to define a basic CLI on client node.

## How it works
```
.
├── client               // Define the client node.
│   ├── clientNode.go    // Create an loop to read stdin stream.
│   └── cmd              // Use Cobra to define the CLI command.
│       └── cmd.go
├── client_node          // Dockerfile - client node.
├── client_node.yml      // docker-compose.yml - client node.
├── conf                 
│   └── app.ini          // System configuration
├── go.mod
├── go.sum
├── LICENSE
├── pb
│   ├── game.pb.go       // The gRPC struct created from game.proto
│   └── game.proto       // Define the protobuf struct.
├── README.md
├── replica              // Replica set configuration
│   ├── config
│   │   └── mongod.yml
│   └── data             // The mongoDB data
│       ├── rs1
│       ├── rs2
│       └── rs3
├── replica_set.js       // MongoDB command used to create replica set.
├── server               // Define the server node.
│   └── serverNode.go    // Create gRPC server.
├── server_node          // Dockerfile - server node.
├── server_node.sh       // Use to run server node after mongoDB started.  
├── server_node.yml      // docker-compose.yml - server node.
└── setting              // Initialize the configuration
    └── setting.go
...
```

### Required

- Golang >= 1.13.7
- docker-compose
- mongo >= 4.2
