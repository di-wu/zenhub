package zenhub

import (
	"fmt"
	"net/http"
)

type IssueData struct {
	IssueNumber  *int        `json:"issue_number,omitempty"`
	RepositoryID *int        `json:"repo_id,omitempty"`
	Estimate     *Estimate   `json:"estimate,omitempty"`
	PlusOnes     []EventType `json:"plus_ones,omitempty"`
	Pipeline     *Pipeline   `json:"pipeline,omitempty"`
	Pipelines    []Pipeline  `json:"pipelines,omitempty"`
	IsEpic       *bool       `json:"is_epic,omitempty"`
}

func (c *Client) GetIssue(repositoryID, issueNumber int) (*IssueData, *http.Response, error) {
	u := fmt.Sprintf("p1/repositories/%d/issues/%d", repositoryID, issueNumber)
	req, err := c.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	issue := new(IssueData)
	resp, err := c.Do(req, issue)
	if err != nil {
		return nil, resp, err
	}
	return issue, resp, nil
}

type IssueEvent struct {
	UserID       *int             `json:"user_id,omitempty"`
	Type         *IssuesEventType `json:"type,omitempty"`
	CreatedAt    *string          `json:"created_at,omitempty"`
	FromEstimate *Estimate        `json:"from_estimate,omitempty"`
	ToEstimate   *Estimate        `json:"to_estimate,omitempty"`
	FromPipeline *Pipeline        `json:"from_pipeline,omitempty"`
	ToPipeline   *Pipeline        `json:"to_pipeline,omitempty"`
	WorkspaceID  *string          `json:"workspace_id,omitempty"`
}

type IssuesEventType string

const (
	EstimateIssue IssuesEventType = "estimateIssue"
	TransferIssue IssuesEventType = "transferIssue"
)

func (c *Client) GetIssueEvents(repositoryID, issueNumber int) (*[]IssueEvent, *http.Response, error) {
	u := fmt.Sprintf("p1/repositories/%d/issues/%d/events", repositoryID, issueNumber)
	req, err := c.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	events := new([]IssueEvent)
	resp, err := c.Do(req, events)
	if err != nil {
		return nil, resp, err
	}
	return events, resp, nil
}

type MoveRequest struct {
	PipelineID string
	Position   Position
}

func (r MoveRequest) toInternal() moveRequest {
	return moveRequest{
		PipelineID: &r.PipelineID,
		Position:   r.Position.value(),
	}
}

type moveRequest struct {
	PipelineID *string     `json:"pipeline_id,omitempty"`
	Position   interface{} `json:"position,omitempty"`
}

type Position struct {
	pos   string
	index int
}

func TopPosition() Position {
	return Position{pos: "top"}
}

func BottomPosition() Position {
	return Position{pos: "bottom"}
}

func NewIndexPosition(value int) Position {
	return Position{index: value}
}

func (p Position) value() interface{} {
	if p.pos != "" {
		return &p.pos
	}
	return &p.index
}

func (c *Client) MoveIssue(workspaceID string, repositoryID, issueNumber int, move MoveRequest) (*http.Response, error) {
	u := fmt.Sprintf("p2/workspaces/%s/repositories/%d/issues/%d/moves", workspaceID, repositoryID, issueNumber)
	req, err := c.NewRequest(http.MethodPost, u, move.toInternal())
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *Client) MoveIssueOld(repositoryID, issueNumber int, move MoveRequest) (*http.Response, error) {
	u := fmt.Sprintf("p1/repositories/%d/issues/%d/moves", repositoryID, issueNumber)
	req, err := c.NewRequest(http.MethodPost, u, move.toInternal())
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

type Estimate struct {
	Value *int `json:"estimate,omitempty"`
}

func (c *Client) SetEstimate(repositoryID, issueNumber, estimate Estimate) (*Estimate, *http.Response, error) {
	u := fmt.Sprintf("p1/repositories/%d/issues%d/estimate", repositoryID, issueNumber)
	req, err := c.NewRequest(http.MethodPut, u, estimate)
	if err != nil {
		return nil, nil, err
	}

	value := new(Estimate)
	resp, err := c.Do(req, value)
	if err != nil {
		return nil, resp, err
	}
	return value, resp, nil
}
