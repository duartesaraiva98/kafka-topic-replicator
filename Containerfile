FROM alpine:3.18.3

ARG TARGETARCH

ARG PATH_TO_EXECUTABLE

COPY kafka-topic-replicator-linux-$TARGETARCH /app/kafka-topic-replicator

CMD ["/app/kafka-topic-replicator"]
