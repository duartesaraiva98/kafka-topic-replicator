FROM alpine:3.18.3

ARG TARGETARCH

ARG PATH_TO_EXECUTABLE

COPY kafka-topic-replicator-linux-$TARGETARCH /opt/kafka-topic-replicator

CMD ["/opt/kafka-topic-replicator"]
