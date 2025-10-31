#!/bin/bash

if [[ -z "$1" ]]; then
  echo "Usage: $0 up|down"
  exit 1
fi

if [[ "$1" == "up" ]]; then
  docker compose -f ~/dev/containers/qbittorrent_nox/docker-compose.yml up -d
  exit 0
elif [[ "$1" == "down" ]]; then
  docker compose -f ~/dev/containers/qbittorrent_nox/docker-compose.yml down
  exit 0
fi
