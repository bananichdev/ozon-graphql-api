#!/bin/bash

psql -U admin <<-EOSQL
    CREATE DATABASE "ozon-graphql-api";
    CREATE DATABASE "ozon-graphql-api-test";
EOSQL