#!/bin/bash

p=`pwd`

arg=$1;

cd $p

case $arg in
  --lint)
    echo "linting goopenov"
    go fmt
    golint
  ;;
  --test)
    echo "testing goopenov"
    go test
  ;;
  --bench)
    echo "testing goopenov"
    go test -bench=.
  ;;
  --build)
    echo "building goopenov"
    env GOOS=linux GOARCH=386 go build -o bin/goopenov # We want our executables to be output to the bin directory
  ;;
  --run)
    echo "Running goopenov"
    cd bin
    ./goopenov
esac

cd $p
