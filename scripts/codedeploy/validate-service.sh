#!/bin/bash

if [[ $(docker inspect --format="{{ .State.Running }}" dp-dd-api-stub) == "false" ]]; then
  exit 1;
fi
