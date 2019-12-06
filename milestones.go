package zenhub

import (
	"fmt"
	"net/http"
	"time"
)

type startDate struct {
	// TODO: use ISO8601 format
	StartDate string `json:"start_date"`
}

func (c *Client) SetMilestoneStartDate(repositoryID, milestoneNumber int, date time.Time) (*time.Time, *http.Response, error) {
	u := fmt.Sprintf("p1/repositories/%d/milestones/%d/start_date", repositoryID, milestoneNumber)
	req, err := c.NewRequest(http.MethodPost, u, startDate{date.Format(time.RFC3339)})

	startDate := new(startDate)
	resp, err := c.Do(req, startDate)
	if err != nil {
		return nil, resp, err
	}
	value, _ := time.Parse(time.RFC3339, startDate.StartDate)
	return &value, resp, nil
}

func (c *Client) GetMilestoneStartDate(repositoryID, milestoneNumber int) (*time.Time, *http.Response, error) {
	u := fmt.Sprintf("p1/repositories/%d/milestones/%d/start_date", repositoryID, milestoneNumber)
	req, err := c.NewRequest(http.MethodGet, u, nil)

	startDate := new(startDate)
	resp, err := c.Do(req, startDate)
	if err != nil {
		return nil, resp, err
	}
	value, _ := time.Parse(time.RFC3339, startDate.StartDate)
	return &value, resp, nil
}
