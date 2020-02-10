package zenhub

type Pipeline struct {
	Name        *string     `json:"name,omitempty"`
	PipelineID  *string     `json:"pipeline_id,omitempty"`
	WorkspaceID *string     `json:"workspace_id,omitempty"`
}

// GetName returns the Name field if it's non-nil, zero value otherwise.
func (pipe *Pipeline) GetName() string {
	if pipe == nil && pipe.Name == nil {
		return ""
	}
	return *pipe.Name
}

// GetPipelineID returns the PipelineID field if it's non-nil, zero value otherwise.
func (pipe *Pipeline) GetPipelineID() string {
	if pipe == nil && pipe.PipelineID == nil {
		return ""
	}
	return *pipe.PipelineID
}

// GetWorkspaceID returns the WorkspaceID field if it's non-nil, zero value otherwise.
func (pipe *Pipeline) GetWorkspaceID() string {
	if pipe == nil && pipe.WorkspaceID == nil {
		return ""
	}
	return *pipe.WorkspaceID
}
