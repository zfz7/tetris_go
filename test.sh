#!/usr/bin/env bash

set -e

cd frontend && CI=true yarn test
cd ../backend && go test -v ./...
cd ../acceptance && POSTGRES_HOST=localhost POSTGRES_PORT=5433 POSTGRES_PASSWORD=tetris_test POSTGRES_USER=tetris_test POSTGRES_DB=tetris_test yarn test