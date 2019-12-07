package zenhub

import (
	"fmt"
	"net/http"
)

type Workspace struct {
	Name         *string `json:"name,omitempty"`
	Description  *string `json:"description,omitempty"`
	ID           *string `json:"id,omitempty"`
	Repositories []int   `json:"repositories,omitempty"`
}

func (c *Client) GetWorkspaces(repositoryID int) (*[]Workspace, *http.Response, error) {
	u := fmt.Sprintf("p2/repositories/%d/workspaces", repositoryID)
	req, err := c.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	workspaces := new([]Workspace)
	resp, err := c.Do(req, workspaces)
	if err != nil {
		return nil, resp, err
	}
	return workspaces, resp, nil
}

type Board struct {
	Pipelines []Pipeline `json:"pipelines"`
}

func (c *Client) GetBoard(repositoryID int, workspaceID string) (*Board, *http.Response, error) {
	u := fmt.Sprintf("p2/workspaces/%s/repositories/%d/board", workspaceID, repositoryID)
	req, err := c.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	board := new(Board)
	resp, err := c.Do(req, board)
	if err != nil {
		return nil, resp, err
	}
	return board, resp, nil
}

func (c *Client) GetBoardOld(repositoryID int) (*Board, *http.Response, error) {
	u := fmt.Sprintf("p1/repositories/%d/board", repositoryID)
	req, err := c.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	board := new(Board)
	resp, err := c.Do(req, board)
	if err != nil {
		return nil, resp, err
	}
	return board, resp, nil
}
