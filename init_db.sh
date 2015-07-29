#!/bin/bash

docker run --name etdb -d -p 5432:5432 -e POSTGRES_USER=etsa -e POSTGRES_PASSWORD=etsaSecret -e POSTGRES_DB=et postgres:latest
