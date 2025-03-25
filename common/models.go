package common

type Repo struct {
	FullName   string `json:"full_name,omitempty"`
	ExternalId string `json:"external_id,omitempty"`
}

type Webhook struct {
	Url string `json:"url"`
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
