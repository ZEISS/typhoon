#!/bin/bash
# This script is executed after the creation of a new project.

sudo apt-get update
sudo apt-get install -y libxml2-dev libxslt1-dev

go install github.com/nats-io/natscli/nats@latest
go install github.com/nats-io/nats-top@latest

# Install the git-hooks via the `ghc` command
# See: https://example/temp
ghc install