FROM alpine:3.18.3

ARG TARGETARCH

COPY ./executables/kafka-topic-replicator-linux-$TARGETARCH /app/kafka-topic-replicator

CMD ["/app/kafka-topic-replicator"]