#!/bin/sh

API_HOST=$(stoml /etc/broquiz/broquiz.toml api.host)
API_PORT=$(stoml /etc/broquiz/broquiz.toml api.port)
echo "Starting BroQuizz on $API_HOST:$API_PORT..."
broquiz-daemon $@