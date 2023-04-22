#!/usr/bin/env bash

script_path="$(dirname $(realpath ${BASH_SOURCE:-$0}))"

OUTPUT="$script_path/prisma/schema.prisma"
MODELS_FILES="$script_path/src/models/**/*.prisma"
DATASOURCE_PATH="$script_path/src/port/datasource.prisma"
GENERATOR_PATH="$script_path/src/port/generator.prisma"

touch $OUTPUT
cat $MODELS_FILES $DATASOURCE_PATH $GENERATOR_PATH > $OUTPUT

