#!/bin/bash
docker run -d --rm --name cnse-redis -p 55013:6379 -p 55014:8001 redis/redis-stack:latest