package zenhub

import (
	"fmt"
	"net/http"
)

func (c *Client) GetAllIssuesForReport(releaseID int) (*[]SimpleIssue, *http.Response, error) {
	u := fmt.Sprintf("p1/reports/release/%d/issues", releaseID)
	req, err := c.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	issues := new([]SimpleIssue)
	resp, err := c.Do(req, issues)
	if err != nil {
		return nil, resp, err
	}
	return issues, resp, nil
}

func (c *Client) UpdateIssueFromReport(releaseID int, update UpdateRequest) (*UpdateResponse, *http.Response, error) {
	u := fmt.Sprintf("p1/reports/release/%d/issues", releaseID)
	req, err := c.NewRequest(http.MethodPatch, u, update)
	if err != nil {
		return nil, nil, err
	}

	updated := new(UpdateResponse)
	resp, err := c.Do(req, updated)
	if err != nil {
		return nil, resp, err
	}
	return updated, resp, nil
}
