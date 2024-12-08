#!/bin/bash

base_dir=$(dirname "$0")

for dir in "$base_dir"/*; do
  echo "checking $dir in $base_dir"
  if [ -d "$dir" ]; then
    echo "checked $dir in $base_dir"
    echo "directory : $dir"

    cd "$dir" || continue

    if [ -f Makefile ]; then
      echo "running 'make up'"
      make up
    else
      echo "there is no makefile in $dir, running manual"
      cp ../.env .env
      dbmate up
    fi

    cd ..
  fi
done
