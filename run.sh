#!/bin/bash

golint ./... \
&& go fmt ./... \
&& make \
&& build/dp-api-spike
