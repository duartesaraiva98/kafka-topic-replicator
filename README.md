# Kafka Topic Replicator

| ![Static Badge](https://img.shields.io/badge/stage-development-green) |

This is a simple tool that helps you with replicating messages in one kafka topic into another. It works in a single or multi cluster setup.

## Configuration

A configuration file needs to be defined and its path needs to be provided through environment variable `CONFIG_FILE`.

### Source

You need to define a source for the replication. The source is the topic you want to make a copy of. Within the source you can configure the consumer that will read the messages fromt the source topic.

```yaml
source:
  topic: source-topic
  consumer:
    group.id: topic-replicator
    bootstrap.servers: localhost:9092
    auto.offset.reset: earliest
```

### Destination

You need to define a destionation of the replication. The destination is the topic you want to copy the messages into. Within the destination you can configure the producer that will emit the messages to the destination topic.

```yaml
destination:
  topic: destination-topic
  producer:
    bootstrap.servers: localhost:9092
```