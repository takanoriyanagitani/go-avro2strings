#!/bin/sh

export ENV_SCHEMA_FILENAME=./sample.d/sample.avsc

#cat sample.d/sample.jsonl | json2avrows > sample.d/input.avro

#export ENV_COLUMN_NAME=active
#export ENV_COLUMN_NAME=name
#export ENV_COLUMN_NAME=id
export ENV_COLUMN_NAME=updated

cat sample.d/input.avro |
	./avro2strings
