default: build

build:
	go build -o ./bin/nats-subscriber ./main.go

docker:
	docker build -t rackhd/nats:client .
