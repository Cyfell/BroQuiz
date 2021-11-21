#!/bin/bash

SCRIPTS_PATH=$( cd "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )
PROJECT_PATH=$(realpath --relative-to=${PWD} ${SCRIPTS_PATH}/..)
API_PATH=${PROJECT_PATH}/api
SOURCE_DIR=${PROJECT_PATH}/cmd

function check {
	echo "Check OpenAPI generation tool..."
	which swagger &> /dev/null || (go get -u github.com/go-swagger/go-swagger/cmd/swagger)
}

function generate {
	check
	echo "Generate OpenAPI specifications from ${SOURCE_DIR} to ${API_PATH}..."
	mkdir -p ${API_PATH}
	swagger generate spec -o ${API_PATH}/swagger.json -w ${SOURCE_DIR} --scan-models
}

function serve {
	check
	echo "Serve OpenAPI specifications at ${API_PATH}..."
	swagger serve -F=swagger ${API_PATH}/swagger.json
}

function help {
	echo "$0: OpenAPI specifications utility"
	echo ""
	echo "Positional arguments:"
	echo "  check:        Check if the OpenAPI tool is preset. If not, install it."
	echo "  generate:     Generate the OpenAPI specifications from source code."
	echo "                Path is ${SWAGGER_FILE_PATH}."
	echo "  serve:        Serve a web data with OpenAPI specifications located"
	echo "                in ${SWAGGER_FILE_PATH}."
	echo "  help:         Display help."
}

while (( "$#" )); do
	case "$1" in
		check) CMD+="check ";;
		generate) CMD+="generate ";;
		serve) CMD+="serve ";;
		help) help;;
		*) CMD+="$1 ";;
	esac
	shift
done

eval $CMD