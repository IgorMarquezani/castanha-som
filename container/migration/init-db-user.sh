#!/bin/bash

set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
  INSERT INTO users values (uuid_generate_v4(), 'root', 'root@gmail.com' 'root123@')
EOSQL
