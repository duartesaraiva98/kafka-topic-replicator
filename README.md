# Kafka Topic Replicator

![Static Badge](https://img.shields.io/badge/stage-development-green)
![Version Badge](https://img.shields.io/github/v/release/duartesaraiva98/kafka-topic-replicator)

This is a simple tool that helps you with replicating messages in one kafka topic into another. It works in a single or multi cluster setup.

## Configuration

A configuration file needs to be defined and its path needs to be provided through environment variable `CONFIG_FILE`. 

### Multi Cluster

Multi cluster configuration should be enabled when the source topic and the destination topic live in different clusters. 

#### Source

You need to define a source for the replication. The source is the topic you want to make a copy of. Within the source you can configure the consumer that will read the messages fromt the source topic.

```yaml
source:
  consumer:
    topic: topic1
    groupId: topic-replicator
    bootstrapServers: localhost:9092
    resetStrategy: earliest
```

#### Destination

You need to define a destionation of the replication. The destination is the topic you want to copy the messages into. Within the destination you can configure the producer that will emit the messages to the destination topic.

```yaml
destination:
  producer:
    topic: topic2
    bootstrapServers: localhost:9092
```

### Single Cluster

Single cluster configuration should be enabled when the source topic and the destination topic live in the same cluster. Provide the bootstrap servers with the key `bootstrapServers` in the configuration file.

#### Source

You need to define a source for the replication. The source is the topic you want to make a copy of. Within the source you can configure the consumer that will read the messages fromt the source topic.

```yaml
source:
  consumer:
    topic: topic1
    groupId: topic-replicator
    resetStrategy: earliest
```

#### Destination

You need to define a destionation of the replication. The destination is the topic you want to copy the messages into. Within the destination you can configure the producer that will emit the messages to the destination topic.

```yaml
destination:
  producer:
    topic: topic2
```