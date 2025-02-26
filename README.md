# circleci-sdk-go

- Official CircleCI sdk for go.
- Client for [CircleCI APIv2](https://circleci.com/docs/api/v2/index.html).
- Generated with [Swagger](https://swagger.io/tools/swagger-codegen/) using:

## How to update
1. Download the newest openapi.json file from [CircleCI APIv2](https://circleci.com/docs/api/v2/index.html) and store it under `circleci/api/openapi.json`.
2. Run the following command:
```
$ ./scripts/generate-sdk.sh
```
