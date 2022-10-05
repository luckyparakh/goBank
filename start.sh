#!/bin/sh

set -e

echo "Start Migration"

/app/migrate -path ./migration -database "$DB_SOURCE" -verbose up

echo "Migration Complete"

exec "$@"