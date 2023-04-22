#!/usr/bin/env bash

script_path="$(dirname $(realpath ${BASH_SOURCE:-$0}))"
MODELS_FILES="$script_path/src/models/*.prisma"
DATASOURCE_PATH="$script_path/src/port/datasource.prisma"
GENERATOR_PATH="$script_path/src/port/generator.prisma"
echo "$MODELS_FILES $DATASOURCE_PATH $GENERATOR_PATH"

cat $MODELS_FILES $DATASOURCE_PATH $GENERATOR_PATH > "$script_path/prisma/schema.prisma"

