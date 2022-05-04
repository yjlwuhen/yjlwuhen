#!/bin/bash

while getopts ":ml" opt
do
  case $opt in
    l)
          CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o animal main.go;;
    m)
          CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o animal main.go;;
    \?)
      echo "Invalid option: -$OPTARG please input option l/m";;
  esac
done