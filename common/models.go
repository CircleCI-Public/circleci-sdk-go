// nolint:revive // introduced before linter
package common

type Repo struct {
	FullName string `json:"full_name,omitempty"`
	// nolint:revive // introduced before linter
	ExternalId string `json:"external_id,omitempty"`
}

type Webhook struct {
	// nolint:revive // introduced before linter
	Url    string `json:"url,omitempty"`
	Sender string `json:"sender,omitempty"`
}

type ConfigSource struct {
	Provider string `json:"provider,omitempty"`
	Repo     Repo   `json:"repo,omitzero"`
	FilePath string `json:"file_path,omitempty"`
}

type CheckoutSource struct {
	Provider string `json:"provider,omitempty"`
	Repo     Repo   `json:"repo,omitzero"`
}

type EventSource struct {
	Provider string  `json:"provider,omitempty"`
	Repo     Repo    `json:"repo,omitzero"`
	Webhook  Webhook `json:"webhook,omitzero"`
}

type PaginatedResponse[T any] struct {
	NextPageToken string `json:"next_page_token"`
	Items         []T    `json:"items"`
}

type VcsInfo struct {
	// nolint:revive // introduced before linter
	VcsUrl        string `json:"vcs_url"`
	Provider      string `json:"provider"`
	DefaultBranch string `json:"default_branch"`
}

type User struct {
	Login string `json:"login"`
}

type Scope struct {
	// nolint:revive // introduced before linter
	Id   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}
