#!/bin/bash

GOARCH=arm \
GOOS=linux \
CGO_ENABLED=1 \
CC=arm-linux-gnueabi-gcc \
    go build