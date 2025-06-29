#!/bin/sh
set -e

echo "Starting the application with mode: $1"

case "$1" in
  all)
    exec ./main all
    ;;
  http)
    exec ./main http
    ;;
  cron)
    exec ./main cron
    ;;
  *)
    echo "Unknown or missing command: '$1'"
    echo "Usage: docker run <image> {all|http|cron}"
    exit 1
    ;;
esac
