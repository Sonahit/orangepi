#!/bin/bash

GOARCH=arm \
GOOS=linux \
CGO_ENABLED=1 \
    go build && ./display