#!/usr/bin/env bash

set -e

cd frontend
yarn build
cd ../
cp -r ./frontend/build ./backend/server
go build ./backend/tetris.go