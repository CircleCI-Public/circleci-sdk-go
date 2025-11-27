package group

import (
	"context"
	"fmt"
	"net/http"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
	"github.com/CircleCI-Public/circleci-sdk-go/common"
)

type Group struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Service struct {
	client *client.Client
}

func NewService(c *client.Client) *Service {
	return &Service{client: c}
}

func (s *Service) Get(ctx context.Context, orgID, groupID string) (_ *Group, err error) {
	var group Group
	_, err = s.client.RequestHelper(ctx, http.MethodGet, fmt.Sprintf("/organizations/%s/groups/%s", orgID, groupID), nil, &group)
	if err != nil {
		return nil, err
	}

	return &group, nil
}

func (s *Service) List(ctx context.Context, orgID string) (_ []Group, err error) {
	var nextPageToken string
	var groupsList []Group
	for {
		var response common.PaginatedResponse[Group]
		_, err = s.client.RequestHelper(ctx, http.MethodGet, fmt.Sprintf("/organizations/%s/groups?page-token=%s", orgID, nextPageToken), nil, &response)
		if err != nil {
			return nil, err
		}

		groupsList = append(groupsList, response.Items...)
		if response.NextPageToken == "" {
			break
		}
		nextPageToken = response.NextPageToken
	}
	return groupsList, nil
}

func (s *Service) Create(ctx context.Context, newGroup Group, orgID string) (_ *Group, err error) {
	var group Group
	_, err = s.client.RequestHelper(ctx, http.MethodPost, fmt.Sprintf("/organizations/%s/groups", orgID), newGroup, &group)
	if err != nil {
		return nil, err
	}

	return &group, nil
}

func (s *Service) Delete(ctx context.Context, orgID, groupID string) (err error) {
	_, err = s.client.RequestHelper(ctx, http.MethodDelete, fmt.Sprintf("/organizations/%s/groups/%s", orgID, groupID), nil, nil)
	return err
}
