FROM rackhd/golang:snap-base

ADD ./build/rootfs/snap-plugin-publisher-nats /snap/auto/snap-plugin-publisher-nats
ADD ./mock-task.yaml /snap/mock-task.yaml
ADD ./snapd-config.yaml /snap/snapd-config.yaml
ADD ./entrypoint.sh /usr/local/bin/entrypoint.sh

ENTRYPOINT ["entrypoint.sh"]

