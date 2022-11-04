#!/bin/bash

FHASH=`md5 $1`

while true; do
  NHASH=`md5 $1`
  if [ "$FHASH" != "$NHASH" ]; then
    ./mdp -file $1
    FHASH=$NHASH
    echo "File changed, reloading..."
  fi
  sleep 5
done