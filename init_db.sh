#!/bin/bash

docker run --name etdb -d -p 5432:5432 -e POSTGRES_USER=${ET_DB_USER} -e POSTGRES_PASSWORD=${ET_DB_PASS} -e POSTGRES_DB=${ET_DB_NAME} postgres:latest
