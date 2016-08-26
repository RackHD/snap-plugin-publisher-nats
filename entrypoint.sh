#!/bin/bash

snapd --config /snap/snapd-config.yaml &

sleep 3
snapctl task create -t /snap/mock-task.yaml

wait

