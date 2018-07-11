#!/bin/sh

# Only use this from Dockerfile

while [ ! -f "$(pwd)/config/production.json" ] || [ -z "$(cat $(pwd)/config/production.json)" ]; do
	echo "Waiting for production.json to be mounted..."
	sleep 10
done

./p2e-background