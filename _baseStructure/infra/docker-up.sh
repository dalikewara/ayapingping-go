#!/bin/sh

docker compose -f docker-compose.yml --project-directory . up -d --force-recreate --build