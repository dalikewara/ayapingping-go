#!/bin/bash

./infra/build.sh
docker compose -f infra/local/docker-compose.yml --project-directory . up -d --force-recreate --build