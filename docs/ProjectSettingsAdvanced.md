# ProjectSettingsAdvanced

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutocancelBuilds** | **bool** | Except for the default branch, cancel running pipelines on a branch when a new pipeline starts on that branch. | [optional] [default to null]
**BuildForkPrs** | **bool** | Run builds for pull requests from forks. | [optional] [default to null]
**BuildPrsOnly** | **bool** | Once enabled, we will only build branches that have associated pull requests open. | [optional] [default to null]
**DisableSsh** | **bool** | When set to true, job re-runs with SSH debugging access will be disabled for the project. | [optional] [default to null]
**ForksReceiveSecretEnvVars** | **bool** | Run builds for forked pull requests with this project&#x27;s configuration, environment variables, and secrets. | [optional] [default to null]
**Oss** | **bool** | Free and Open Source. Enabling this grants additional credits, and lets others see your builds, both through the web UI and the API. | [optional] [default to null]
**SetGithubStatus** | **bool** | Report the status of every pushed commit to GitHub&#x27;s status API. Updates reported per job. | [optional] [default to null]
**SetupWorkflows** | **bool** | Enabling allows you to conditionally trigger configurations outside of the primary &#x60;.circleci&#x60; parent directory. | [optional] [default to null]
**WriteSettingsRequiresAdmin** | **bool** | Whether updating these settings requires a user to be an organization administrator. When disabled, updating settings can be done by any member. | [optional] [default to null]
**PrOnlyBranchOverrides** | **[]string** | This field is used in conjunction with the &#x60;build_prs_only&#x60;, it allows you to specify a list of branches that will always triger a build. The value passed will overwrite the existing value. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

