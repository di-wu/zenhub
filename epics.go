package zenhub

import (
	"fmt"
	"net/http"
)

type EpicIssue struct {
	IssueNumber  *int    `json:"issue_number,omitempty"`
	RepositoryID *int    `json:"repo_id,omitempty"`
	IssueURL     *string `json:"issue_url,omitempty"`
}

func (c *Client) GetEpics(repositoryID int) (*[]EpicIssue, *http.Response, error) {
	u := fmt.Sprintf("p1/repositories/%d/epics", repositoryID)
	req, err := c.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	issues := new([]EpicIssue)
	resp, err := c.Do(req, issues)
	if err != nil {
		return nil, resp, err
	}
	return issues, resp, nil
}

type EpicData struct {
	TotalEpicEstimates *Estimate   `json:"total_epic_estimates,omitempty"`
	Estimate           *Estimate   `json:"estimate,omitempty"`
	Pipeline           *Pipeline   `json:"pipeline,omitempty"`
	Pipelines          []Pipeline  `json:"pipelines,omitempty"`
	Issues             []IssueData `json:"issues,omitempty"`
}

func (c *Client) GetEpic(repositoryID, epicID int) (*EpicData, *http.Response, error) {
	u := fmt.Sprintf("p1/repositories/%d/epics/%d", repositoryID, epicID)
	req, err := c.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	epic := new(EpicData)
	resp, err := c.Do(req, epic)
	if err != nil {
		return nil, resp, err
	}
	return epic, resp, nil
}

func (c *Client) ConvertEpicToIssue(repositoryID, issueNumber int) (*http.Response, error) {
	u := fmt.Sprintf("p1/repositories/%d/epics/%d/convert_to_issue", repositoryID, issueNumber)
	req, err := c.NewRequest(http.MethodPost, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

type ConvertIssueRequest struct {
	Issues []SimpleIssue `json:"issues"`
}

type SimpleIssue struct {
	RepositoryID int `json:"repo_id"`
	IssueNumber  int `json:"issue_number"`
}

func (c *Client) ConvertIssueToEpic(repositoryID, issueNumber int, issues ConvertIssueRequest) (*http.Response, error) {
	u := fmt.Sprintf("p1/repositories/%d/issues/%d/convert_to_epic", repositoryID, issueNumber)
	req, err := c.NewRequest(http.MethodPost, u, issues)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

type UpdateRequest struct {
	RemoveIssues []SimpleIssue `json:"remove_issues"`
	AddIssues    []SimpleIssue `json:"add_issues"`
}

type UpdateResponse struct {
	RemovedIssues []SimpleIssue `json:"removed_issues,omitempty"`
	AddedIssues   []SimpleIssue `json:"added_issues,omitempty"`
}

func (c *Client) UpdateIssuesToEpic(repositoryID, issueNumber int, update UpdateRequest) (*UpdateResponse, *http.Response, error) {
	u := fmt.Sprintf("p1/repositories/%d/epics/%d/update_issues", repositoryID, issueNumber)
	req, err := c.NewRequest(http.MethodPost, u, update)
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
