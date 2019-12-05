package zenhub

type Event struct {
	Type         Type   `json:"type"`
	GitHubURL    string `json:"url"`
	Organization string `json:"organization"`
	Repository   string `json:"repo"`
	UserName     string `json:"user_name"`
	IssueNumber  string `json:"issue_number"`
	IssueTitle   string `json:"issue_title"`
}

type Type string

const (
	IssueTransfer      Type = "issue_transfer"
	EstimateSet        Type = "estimate_set"
	EstimateCleared    Type = "estimate_cleared"
	IssueReprioritized Type = "issue_reprioritized"
)

type IssueTransferEvent struct {
	Event
	ToPipelineName   string `json:"to_pipeline_name"`
	FromPipelineName string `json:"from_pipeline_name"`
}

type EstimateSetEvent struct {
	Event
	Estimate string `json:"estimate"`
}

type IssueReprioritizedEvent struct {
	Event
	ToPipelineName string `json:"to_pipeline_name"`
	FromPosition   string `json:"from_position"`
	ToPosition     string `json:"to_position"`
}
