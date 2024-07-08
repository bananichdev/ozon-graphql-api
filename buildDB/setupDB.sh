#!/bin/bash

psql -U admin <<-EOSQL
    CREATE DATABASE "ozon-graphql-api";
EOSQL