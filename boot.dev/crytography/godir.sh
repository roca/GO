#!/bin/sh -f

DIR=$1
README_TEXT=$2

if [ -z $DIR ]; then
  echo "Usage: godir <dir>"
  exit 1
fi

echo "Intializing Go directory $DIR"

mkdir $DIR
$(cd $DIR && go mod init $DIR && touch main.go && touch main_test.go)
echo "package main" > $DIR/main.go
echo "package main" > $DIR/main_test.go
if [ -z $README_TEXT ]; then
  README_TEXT="# This is a Go project"
fi
echo  "# $DIR: $README_TEXT" > $DIR/README.md
go work use $DIR
