package zenhub

import (
	"fmt"
	"net/http"
)

type Dependencies struct {
	Dependencies []Dependency `json:"dependencies"`
}

type Dependency struct {
	Blocking SimpleIssue `json:"blocking"`
	Blocked  SimpleIssue `json:"blocked"`
}

func (c *Client) GetDependencies(repositoryID int) (*Dependencies, *http.Response, error) {
	u := fmt.Sprintf("p1/repositories/%d/dependencies", repositoryID)
	req, err := c.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	dependencies := new(Dependencies)
	resp, err := c.Do(req, dependencies)
	if err != nil {
		return nil, resp, err
	}
	return dependencies, resp, nil
}

func (c *Client) CreateDependency(dependency Dependency) (*Dependency, *http.Response, error) {
	u := fmt.Sprintf("p1/dependencies")
	req, err := c.NewRequest(http.MethodPost, u, dependency)
	if err != nil {
		return nil, nil, err
	}

	dep := new(Dependency)
	resp, err := c.Do(req, dep)
	if err != nil {
		return nil, resp, err
	}
	return dep, resp, nil
}

func (c *Client) RemoveDependency(dependency Dependency) (*http.Response, error) {
	u := fmt.Sprintf("p1/dependencies")
	req, err := c.NewRequest(http.MethodDelete, u, dependency)
	if err != nil {
		return nil, err
	}

	dep := new(Dependency)
	resp, err := c.Do(req, dep)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
