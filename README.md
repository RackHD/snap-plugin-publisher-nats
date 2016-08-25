# snap-plugin-publisher-nats  

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

