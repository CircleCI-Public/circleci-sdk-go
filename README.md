# CircleCI SDK GO


|RESOURCE | GET | CREATE | LIST | UPDATE | DELETE |
|---------|-----|--------|------|--------|--------|
|_project_|stable|stable|❌|❌|❌|
|_project settings_|stable|❌|❌|stable|❌|
|_pipeline definition_|stable|stable|stable|stable|stable|
|_trigger_|stable|stable|stable|stable|stable|
|_context_|stable|stable|stable|❌|stable|
|_context restriction_|❌|stable|stable|❌|stable|
|_env var_|❌|stable|stable|❌|stable|

# Development

This repository makes use [Task](https://taskfile.dev/#/). It may be installed (on MacOS) with:
```
$ brew install go-task/tap/go-task
```

See the full list of available tasks by running `task -l`, or, see the [Taskfile.yml](./Taskfile.yml) script.

```sh
task lint
task fmt
task generate

# Run all the tests
task test
# Run the tests for one package
task test -- ./client/...
# Run all the quick tests
task test -- -short ./...
