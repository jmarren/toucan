#!/bin/bash

DB_URL=postgres://postgres:password@localhost:5434/postgres?sslmode=disable

migrate -database $DB_URL -path ./db/migrations up

