package zenhub

type Pipeline struct {
	Name        *string     `json:"name,omitempty"`
	PipelineID  *string     `json:"pipeline_id,omitempty"`
	WorkspaceID *string     `json:"workspace_id,omitempty"`
	Issues      []IssueData `json:"issues,omitempty"`
}
