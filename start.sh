#!/usr/bin/env bash

set -e

./build.sh

if [ "$1" == 'test' ]; then
	POSTGRES_HOST=localhost SERVER_PORT=8444 POSTGRES_PORT=5433 POSTGRES_PASSWORD=tetris_test POSTGRES_USER=tetris_test POSTGRES_DB=tetris_test ./tetris
else
	POSTGRES_HOST=localhost SERVER_PORT=8080 POSTGRES_PORT=5434 POSTGRES_PASSWORD=tetris POSTGRES_USER=tetris POSTGRES_DB=tetris ./tetris
fi
