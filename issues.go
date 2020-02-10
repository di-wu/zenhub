package zenhub

import (
	"fmt"
	"net/http"
)

type IssueData struct {
	Estimate  *Estimate   `json:"estimate,omitempty"`
	PlusOnes  []*PlusOne  `json:"plus_ones,omitempty"`
	Pipeline  *Pipeline   `json:"pipeline,omitempty"`
	Pipelines []*Pipeline `json:"pipelines,omitempty"`
	IsEpic    *bool       `json:"is_epic,omitempty"`
}

// GetEstimate returns the Estimate.Value field if it's non-nil, zero value otherwise.
func (data *IssueData) GetEstimate() int {
	if data == nil || data.Estimate == nil || data.Estimate.Value == nil {
		return 0
	}
	return *data.Estimate.Value
}

type PlusOne struct {
	UserID    *int    `json:"user_id"`
	CreatedAt *string `json:"created_at"`
}

// GetUserID returns the UserID field if it's non-nil, zero value otherwise.
func (plus *PlusOne) GetUserID() int {
	if plus == nil && plus.UserID == nil {
		return 0
	}
	return *plus.UserID
}

// GetCreatedAt returns the CreatedAt field if it's non-nil, zero value otherwise.
func (plus *PlusOne) GetCreatedAt() string {
	if plus == nil && plus.UserID == nil {
		return ""
	}
	return *plus.CreatedAt
}

// GetIsEpic returns the IsEpic field if it's non-nil, zero value otherwise.
func (data *IssueData) GetIsEpic() bool {
	if data == nil && data.IsEpic == nil {
		return false
	}
	return *data.IsEpic
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

// GetUserID returns the UserID field if it's non-nil, zero value otherwise.
func (event *IssueEvent) GetUserID() int {
	if event == nil || event.UserID == nil {
		return 0
	}
	return *event.UserID
}

// GetType returns the Type field if it's non-nil, zero value otherwise.
func (event *IssueEvent) GetType() IssuesEventType {
	if event == nil || event.Type == nil {
		return ""
	}
	return *event.Type
}

// GetCreatedAt returns the CreatedAt field if it's non-nil, zero value otherwise.
func (event *IssueEvent) GetCreatedAt() string {
	if event == nil || event.CreatedAt == nil {
		return ""
	}
	return *event.CreatedAt
}

// GetFromEstimate returns the FromEstimate field if it's non-nil, zero value otherwise.
func (event *IssueEvent) GetFromEstimate() int {
	if event == nil {
		return 0
	}
	return event.FromEstimate.GetValue()
}

// GetToEstimate returns the ToEstimate field if it's non-nil, zero value otherwise.
func (event *IssueEvent) GetToEstimate() int {
	if event == nil {
		return 0
	}
	return event.ToEstimate.GetValue()
}

// GetWorkspaceID returns the WorkspaceID field if it's non-nil, zero value otherwise.
func (event *IssueEvent) GetWorkspaceID() string {
	if event == nil && event.WorkspaceID == nil {
		return ""
	}
	return *event.WorkspaceID
}

type IssuesEventType string

const (
	EstimateIssue IssuesEventType = "estimateIssue"
	TransferIssue IssuesEventType = "transferIssue"
)

func (c *Client) GetIssueEvents(repositoryID, issueNumber int) ([]*IssueEvent, *http.Response, error) {
	u := fmt.Sprintf("p1/repositories/%d/issues/%d/events", repositoryID, issueNumber)
	req, err := c.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var events []*IssueEvent
	resp, err := c.Do(req, &events)
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

// GetValue returns the Value field if it's non-nil, zero value otherwise.
func (e *Estimate) GetValue() int {
	if e == nil || e.Value == nil {
		return 0
	}
	return *e.Value
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
