package zenhub

import (
	"fmt"
	"net/http"
)

type Report struct {
	ReleaseID      *string `json:"release_id"`
	Title          *string `json:"title"`
	Description    *string `json:"description"`
	StartDate      *string `json:"start_date"`
	DesiredEndDate *string `json:"desired_end_date"`
	CreatedAt      *string `json:"created_at"`
	ClosedAt       *string `json:"closed_at"`
	State          *string `json:"state"`
	Repositories   []int   `json:"repositories"`
}

func (c *Client) CreateReleaseReport(repositoryID int, report Report) (*Report, *http.Response, error) {
	u := fmt.Sprintf("p1/repositories/%d/reports/release", repositoryID)
	req, err := c.NewRequest(http.MethodPost, u, report)
	if err != nil {
		return nil, nil, err
	}

	rep := new(Report)
	resp, err := c.Do(req, rep)
	if err != nil {
		return nil, resp, err
	}
	return rep, resp, nil
}

func (c *Client) GetReleaseReport(releaseID int) (*Report, *http.Response, error) {
	u := fmt.Sprintf("p1/reports/release/%d", releaseID)
	req, err := c.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	report := new(Report)
	resp, err := c.Do(req, report)
	if err != nil {
		return nil, resp, err
	}
	return report, resp, nil
}

func (c *Client) GetReleaseReports(repositoryID int) (*[]Report, *http.Response, error) {
	u := fmt.Sprintf("p1/repositories/%d/reports/releases", repositoryID)
	req, err := c.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	reports := new([]Report)
	resp, err := c.Do(req, reports)
	if err != nil {
		return nil, resp, err
	}
	return reports, resp, nil
}

func (c *Client) EditReleaseReport(releaseID int, report Report) (*Report, *http.Response, error) {
	u := fmt.Sprintf("p1/reports/release/%d", releaseID)
	req, err := c.NewRequest(http.MethodPatch, u, report)
	if err != nil {
		return nil, nil, err
	}

	rep := new(Report)
	resp, err := c.Do(req, report)
	if err != nil {
		return nil, resp, err
	}
	return rep, resp, nil
}

func (c *Client) AddRepositoryToReleaseReport(releaseID, repositoryID int) (*http.Response, error) {
	u := fmt.Sprintf("p1/reports/release/%d/repository/%d", releaseID, repositoryID)
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

func (c *Client) RemoveRepositoryToReleaseReport(releaseID, repositoryID int) (*http.Response, error) {
	u := fmt.Sprintf("p1/reports/release/%d/repository/%d", releaseID, repositoryID)
	req, err := c.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
