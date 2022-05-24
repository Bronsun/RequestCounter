#!/bin/bash

export PORT
export ADDR

envsubst '${PORT} ${ADDR}' < /nginx.conf.template > /etc/nginx/nginx.conf

exec /docker-entrypoint.sh "$@"