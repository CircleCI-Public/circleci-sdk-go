#!/usr/bin/env bash

swagger-codegen generate -l go -i circleci/api/openapi.json -o circleci -c swagger/config.json
