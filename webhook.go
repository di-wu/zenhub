package zenhub

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	ErrNoTypesSpecified  = errors.New("no type specified to parse")
	ErrInvalidHTTPMethod = errors.New("invalid http method")
	ErrParsingPayload    = errors.New("error parsing payload")
	ErrEventNotFound     = errors.New("event not defined to be parsed")
)

// Webhook instance that contain the method to process incoming events.
type Webhook struct{}

// Parse parses and verifies the specified types and returns an event object or an error.
func (h Webhook) Parse(r *http.Request, types ...EventType) (interface{}, error) {
	if len(types) == 0 {
		return nil, ErrNoTypesSpecified
	}
	if r.Method != http.MethodPost {
		return nil, ErrInvalidHTTPMethod
	}

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil || len(payload) == 0 {
		return nil, ErrParsingPayload
	}

	q, err := url.ParseQuery(string(payload))
	if err != nil {
		return nil, err
	}
	values := make(map[string]string)
	for key, _ := range q {
		values[key] = q.Get(key)
	}

	js, err := json.Marshal(values)
	if err != nil {
		return nil, err
	}

	var event Event
	if err := json.Unmarshal(js, &event); err != nil {
		return nil, err
	}

	var found bool
	for _, typ := range types {
		if typ == event.Type {
			found = true
			break
		}
	}
	if !found {
		return nil, ErrEventNotFound
	}

	switch event.Type {
	case IssueTransfer:
		var pl IssueTransferEvent
		err = json.Unmarshal(js, &pl)
		return pl, err
	case EstimateSet:
		var pl EstimateSetEvent
		err = json.Unmarshal(js, &pl)
		return pl, err
	case EstimateCleared:
		return EstimateClearedEvent{event}, nil
	case IssueReprioritized:
		var pl IssueReprioritizedEvent
		err = json.Unmarshal(js, &pl)
		return pl, err
	default:
		return nil, fmt.Errorf("unknown type %s", event.Type)
	}
}

type Event struct {
	Type         EventType `json:"type"`
	GitHubURL    string    `json:"url"`
	Organization string    `json:"organization"`
	Repository   string    `json:"repo"`
	UserName     string    `json:"user_name"`
	IssueNumber  string    `json:"issue_number"`
	IssueTitle   string    `json:"issue_title"`
}

type EventType string

const (
	IssueTransfer      EventType = "issue_transfer"
	EstimateSet        EventType = "estimate_set"
	EstimateCleared    EventType = "estimate_cleared"
	IssueReprioritized EventType = "issue_reprioritized"
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

type EstimateClearedEvent struct {
	Event
}

type IssueReprioritizedEvent struct {
	Event
	ToPipelineName string `json:"to_pipeline_name"`
	FromPosition   string `json:"from_position"`
	ToPosition     string `json:"to_position"`
}
