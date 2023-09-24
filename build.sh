#!/bin/bash
set -ex

OS=$1
ARCH=$2
VERSION=$3

go build -C src -o target/kafka-topic-replicator-$VERSION-$OS-$ARCH