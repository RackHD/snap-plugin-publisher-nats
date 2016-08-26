# snap-plugin-publisher-nats  

## Using Docker  
The following steps start and run 3 containers: (1) Snap, with this Nats publisher plugin, (2) a Nats server for the plugin to publish to, and (3) a Nats client to listen to the server  

1. `make`
2. `make docker`
3. `docker-compose up`


If you want to run Snap with this plugin directly on a host, and not in a container, use the remainder of this README.  


## Starting a Nats server  
1. `go get github.com/nats-io/gnatsd`  
2. `cd $GOPATH/src/github.com/nats-io/gnatsd`  
3. `./gnatsd -a 127.0.0.1`  

## Build this plugin 
1. `cd $GOPATH/src/github.com/skunkworxs/snap-plugin-publisher-nats`
2. `make`   

## Adding it to Snap  
1. Have Snap running  
2. Load a Collector and Processor Plugin  
3. `cd $GOPATH/src/github.com/skunkworxs/snap-plugin-publisher-nats`  
4. `snapctl plugin load build/rootfs/snap-plugin-publisher-nats`  

## Running a Snap task  
1. `cd $GOPATH/src/github.com/skunkworxs/snap-plugin-publisher-nats`  
2. `snapctl task create -t mock-file.yaml`  

## Listening using the Nats listener
1. `cd $GOPATH/src/github.com/skunkworxs/snap-plugin-publisher-nats/examples/nats-client`  
2. `make`  
3. `./bin/nats-subscriber`  

