#!/bin/bash


SERVICE_NAME=depeche_backend
docker exec $SERVICE_NAME fuser -k /depeche-backend/main
docker exec $SERVICE_NAME make build