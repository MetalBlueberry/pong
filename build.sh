#!/bin/bash

GOOS=js GOARCH=wasm go build -o ./docs/pong.wasm   ./cmd/pong/