#!/bin/bash

if [[ $(docker inspect --format="{{ .State.Running}}" frontend-renderer) == "false" ]]; then
  exit 1;
fi
