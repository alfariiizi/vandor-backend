#!/bin/bash
MODULE=$1

if ! command -v air &> /dev/null; then
	read -p "'air' is not installed. Install it? [Y/n] " choice
	if [[ $choice != "n" && $choice != "N" ]]; then
		go install github.com/air-verse/air@latest
	else
		echo "Exiting..."
		exit 1
	fi
fi

echo "Watching module: $MODULE"
air -c .air.toml $MODULE
