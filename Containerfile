FROM alpine:3.18.3

ARG TARGETARCH

COPY src/target/kafka-topic-replicator-linux-$TARGETARCH /app/kafka-topic-replicator

CMD ["kafka-topic-replicator"]