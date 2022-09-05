#!/usr/bin/env bash

set -e

cd frontend && CI=true yarn test
cd ../backend && go test
cd ../acceptance && yarn test