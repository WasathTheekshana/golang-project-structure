#!/bin/bash

# Connect to PostgreSQL and create the uuid-ossp extension
docker exec -i golang-proj-struct psql -U admin -d golang_proj <<EOF
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
EOF
