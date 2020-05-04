#!/usr/bin/env bash

export GO111MODULE=on

go build -o output/timeline

cp -R templates output/

mkdir output/logs
