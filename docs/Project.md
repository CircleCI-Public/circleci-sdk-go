# Project

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slug** | **string** | Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. For projects that use GitLab or GitHub App, use &#x60;circleci&#x60; as the &#x60;vcs-slug&#x60;, replace &#x60;org-name&#x60; with the organization ID (found in Organization Settings), and replace &#x60;repo-name&#x60; with the project ID (found in Project Settings). | [default to null]
**Name** | **string** | The name of the project | [default to null]
**Id** | **string** |  | [default to null]
**OrganizationName** | **string** | The name of the organization the project belongs to | [default to null]
**OrganizationSlug** | **string** | The slug of the organization the project belongs to | [default to null]
**OrganizationId** | **string** | The id of the organization the project belongs to | [default to null]
**VcsInfo** | [***ProjectVcsInfo**](Project_vcs_info.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

