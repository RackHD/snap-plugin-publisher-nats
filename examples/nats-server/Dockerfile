FROM debian:wheezy

ENV NATS_V v0.9.4

RUN mkdir /nats \

  && apt-get update && apt-get install -y --no-install-recommends ca-certificates curl unzip && rm -rf /var/lib/apt/lists/* \

  && cd /nats \
  && curl -OLs https://github.com/nats-io/gnatsd/releases/download/$NATS_V/gnatsd-$NATS_V-linux-amd64.zip \

  && unzip gnatsd-$NATS_V-linux-amd64.zip \
  && mv    gnatsd-$NATS_V-linux-amd64/gnatsd /usr/local/bin/ \

  && cd / \
  && rm -rf /nats

ENTRYPOINT ["gnatsd"]
