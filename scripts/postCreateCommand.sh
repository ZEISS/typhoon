#!/bin/bash
# This script is executed after the creation of a new project.

sudo apt-get update
sudo apt-get install libxml2-dev libxslt1-dev

go install github.com/air-verse/air@latest