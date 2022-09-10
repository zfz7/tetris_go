#!/usr/bin/env bash

set -e

cd frontend
yarn build
cd ../
cp -r ./frontend/build ./backend/server

if [ "$1" == 'prod' ]; then
  GOOS=linux GOARCH=amd64 go build ./backend/tetris.go
else
  go build ./backend/tetris.go
fi
