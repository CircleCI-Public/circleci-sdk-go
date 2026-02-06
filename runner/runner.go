package runner

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
)

// Runner represents a self-hosted runner instance.
type Runner struct {
	Name           string `json:"name,omitempty"`
	Hostname       string `json:"hostname,omitempty"`
	IP             string `json:"ip,omitempty"`
	Version        string `json:"version,omitempty"`
	Status         string `json:"status,omitempty"`
	ResourceClass  string `json:"resource_class,omitempty"`
	FirstConnected string `json:"first_connected,omitempty"`
	LastConnected  string `json:"last_connected,omitempty"`
	LastUsed       string `json:"last_used,omitempty"`
}

// ResourceClass represents a runner resource class configuration.
type ResourceClass struct {
	//nolint:revive
	Id            string `json:"id,omitempty"`
	ResourceClass string `json:"resource_class,omitempty"`
	Description   string `json:"description,omitempty"`
}

// Token represents a runner authentication token.
type Token struct {
	//nolint:revive
	Id            string `json:"id,omitempty"`
	Nickname      string `json:"nickname,omitempty"`
	ResourceClass string `json:"resource_class,omitempty"`
	Token         string `json:"token,omitempty"` // Only returned on create
	CreatedAt     string `json:"created_at,omitempty"`
}

// CreateResourceClassRequest contains the parameters for creating a resource class.
type CreateResourceClassRequest struct {
	ResourceClass string `json:"resource_class"`
	Description   string `json:"description,omitempty"`
}

// CreateTokenRequest contains the parameters for creating a runner token.
type CreateTokenRequest struct {
	ResourceClass string `json:"resource_class"`
	Nickname      string `json:"nickname"`
}

// ListRunnersParams contains filtering parameters for listing runners.
type ListRunnersParams struct {
	ResourceClass string
	Namespace     string
	OrgID         string
}

// UnclaimedTaskCount represents the count of unclaimed tasks for a resource class.
type UnclaimedTaskCount struct {
	UnclaimedTaskCount int `json:"unclaimed_task_count"`
}

// RunningTaskCount represents the count of running tasks for a resource class.
type RunningTaskCount struct {
	RunningRunnerTasks int `json:"running_runner_tasks"`
}

// ClaimTaskRequest contains the parameters for claiming a task.
type ClaimTaskRequest struct {
	Hostname string `json:"hostname"`
	IP       string `json:"ip"`
	Name     string `json:"name"`
	Version  string `json:"version"`
}

// ClaimTaskResponse contains the response from claiming a task.
type ClaimTaskResponse struct {
	AgentVersion  string `json:"agent_version,omitempty"`
	Allocation    string `json:"allocation,omitempty"`
	IsEnterprise  bool   `json:"is_enterprise,omitempty"`
	ResourceClass string `json:"resource_class,omitempty"`
	TaskToken     string `json:"task_token,omitempty"`
	Warning       string `json:"warning,omitempty"`
}

// UnclaimTaskRequest contains the parameters for unclaiming a task.
type UnclaimTaskRequest struct {
	//nolint:revive
	TaskId    string `json:"task_id"`
	TaskToken string `json:"task_token"`
}

// Service provides methods for interacting with the CircleCI Runner Admin API v3.
type Service struct {
	client *client.Client
}

// NewService creates a new Service instance.
func NewService(c *client.Client) *Service {
	return &Service{client: c}
}

// ListRunners returns a list of runners filtered by the provided parameters.
// At least one filter parameter should be provided.
func (s *Service) ListRunners(ctx context.Context, params ListRunnersParams) ([]Runner, error) {
	values := url.Values{}
	if params.ResourceClass != "" {
		values.Add("resource-class", params.ResourceClass)
	}
	if params.Namespace != "" {
		values.Add("namespace", params.Namespace)
	}
	if params.OrgID != "" {
		values.Add("org-id", params.OrgID)
	}

	query := ""
	if len(values) > 0 {
		query = "?" + values.Encode()
	}

	var runners []Runner
	_, err := s.client.RequestHelper(ctx, http.MethodGet, "/api/v3/runner"+query, nil, &runners)
	if err != nil {
		return nil, err
	}

	return runners, nil
}

// ListResourceClasses returns a list of resource classes filtered by namespace and/or organization ID.
// At least one filter parameter should be provided.
func (s *Service) ListResourceClasses(ctx context.Context, namespace, orgID string) ([]ResourceClass, error) {
	values := url.Values{}
	if namespace != "" {
		values.Add("namespace", namespace)
	}
	if orgID != "" {
		values.Add("org-id", orgID)
	}

	query := ""
	if len(values) > 0 {
		query = "?" + values.Encode()
	}

	var resourceClasses []ResourceClass
	_, err := s.client.RequestHelper(ctx, http.MethodGet, "/api/v3/runner/resource"+query, nil, &resourceClasses)
	if err != nil {
		return nil, err
	}

	return resourceClasses, nil
}

// CreateResourceClass creates a new runner resource class.
func (s *Service) CreateResourceClass(ctx context.Context, req CreateResourceClassRequest) (*ResourceClass, error) {
	var resourceClass ResourceClass
	_, err := s.client.RequestHelper(ctx, http.MethodPost, "/api/v3/runner/resource", req, &resourceClass)
	if err != nil {
		return nil, err
	}

	return &resourceClass, nil
}

// DeleteResourceClass deletes a resource class by ID.
// If force is true, the resource class will be deleted even if it has associated tokens.
func (s *Service) DeleteResourceClass(ctx context.Context, id string, force bool) error {
	path := fmt.Sprintf("/api/v3/runner/resource/%s", id)
	if force {
		path += "/force"
	}

	_, err := s.client.RequestHelper(ctx, http.MethodDelete, path, nil, nil)
	return err
}

// ListTokens returns a list of tokens for a specific resource class.
func (s *Service) ListTokens(ctx context.Context, resourceClass string) ([]Token, error) {
	values := url.Values{}
	if resourceClass != "" {
		values.Add("resource-class", resourceClass)
	}

	query := ""
	if len(values) > 0 {
		query = "?" + values.Encode()
	}

	var tokens []Token
	_, err := s.client.RequestHelper(ctx, http.MethodGet, "/api/v3/runner/token"+query, nil, &tokens)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

// CreateToken creates a new runner token.
// The token value is only returned in the response and cannot be retrieved later.
func (s *Service) CreateToken(ctx context.Context, req CreateTokenRequest) (*Token, error) {
	var token Token
	_, err := s.client.RequestHelper(ctx, http.MethodPost, "/api/v3/runner/token", req, &token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

// DeleteToken deletes a runner token by ID.
func (s *Service) DeleteToken(ctx context.Context, id string) error {
	_, err := s.client.RequestHelper(ctx, http.MethodDelete, fmt.Sprintf("/api/v3/runner/token/%s", id), nil, nil)
	return err
}

// GetUnclaimedTaskCount returns the number of unclaimed tasks for a resource class.
func (s *Service) GetUnclaimedTaskCount(ctx context.Context, resourceClass string) (*UnclaimedTaskCount, error) {
	values := url.Values{}
	if resourceClass != "" {
		values.Add("resource-class", resourceClass)
	}

	query := ""
	if len(values) > 0 {
		query = "?" + values.Encode()
	}

	var count UnclaimedTaskCount
	_, err := s.client.RequestHelper(ctx, http.MethodGet, "/api/v3/runner/tasks"+query, nil, &count)
	if err != nil {
		return nil, err
	}

	return &count, nil
}

// GetRunningTaskCount returns the number of running tasks for a resource class.
func (s *Service) GetRunningTaskCount(ctx context.Context, resourceClass string) (*RunningTaskCount, error) {
	values := url.Values{}
	if resourceClass != "" {
		values.Add("resource-class", resourceClass)
	}

	query := ""
	if len(values) > 0 {
		query = "?" + values.Encode()
	}

	var count RunningTaskCount
	_, err := s.client.RequestHelper(ctx, http.MethodGet, "/api/v3/runner/tasks/running"+query, nil, &count)
	if err != nil {
		return nil, err
	}

	return &count, nil
}

// ClaimTask attempts to claim a task for execution by a runner.
// Returns task details if a task is available, or may return a 204 status if no tasks are available.
func (s *Service) ClaimTask(ctx context.Context, req ClaimTaskRequest) (*ClaimTaskResponse, error) {
	var response ClaimTaskResponse
	_, err := s.client.RequestHelper(ctx, http.MethodPost, "/api/v3/runner/claim", req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// UnclaimTask releases a previously claimed task back to the queue.
func (s *Service) UnclaimTask(ctx context.Context, req UnclaimTaskRequest) error {
	_, err := s.client.RequestHelper(ctx, http.MethodPost, "/api/v3/runner/unclaim", req, nil)
	return err
}
