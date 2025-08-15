#!/bin/sh
host="$1"
shift
cmd="$@"

# Wait until Postgres port is open
while ! nc -z "$host" 5432; do
  echo "Waiting for Postgres at $host..."
  sleep 2
done

exec "$@"
