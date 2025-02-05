package ilert

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
)

// Incident definition
type Incident struct {
	ID                 int64                  `json:"id"`
	Summary            string                 `json:"summary"`
	Details            string                 `json:"details"`
	ReportTime         string                 `json:"reportTime"` // Date time string in ISO format
	ResolvedOn         string                 `json:"resolvedOn"` // Date time string in ISO format
	Status             string                 `json:"status"`
	AlertSource        *AlertSource           `json:"alertSource,omitempty"`
	Priority           string                 `json:"priority"`
	IncidentKey        string                 `json:"incidentKey"`
	AssignedTo         *User                  `json:"assignedTo,omitempty"`
	NextEscalation     string                 `json:"nextEscalation"` // Date time string in ISO format
	CallRoutingNumber  *CallRoutingNumber     `json:"callRoutingNumber,omitempty"`
	AcknowledgedBy     *User                  `json:"acknowledgedBy,omitempty"`
	AcknowledgedByType string                 `json:"acknowledgedByType,omitempty"`
	ResolvedBy         *User                  `json:"resolvedBy,omitempty"`
	ResolvedByType     string                 `json:"resolvedByType,omitempty"`
	Images             []IncidentImage        `json:"images,omitempty"`
	Links              []IncidentLink         `json:"links,omitempty"`
	CustomDetails      map[string]interface{} `json:"customDetails,omitempty"`
}

// IncidentImage represents event image
type IncidentImage struct {
	Src  string `json:"src"`
	Href string `json:"href"`
	Alt  string `json:"alt"`
}

// IncidentLink represents event link
type IncidentLink struct {
	Text string `json:"text"`
	Href string `json:"href"`
}

// IncidentComment definition
type IncidentComment struct {
	ID             string `json:"id"`
	Content        string `json:"content"`
	Creator        *User  `json:"creator"`
	TriggerType    string `json:"triggerType"`
	ResolveComment bool   `json:"resolveComment"`
	Created        string `json:"created"`
	Updated        string `json:"updated"`
}

// CallRoutingNumber definition
type CallRoutingNumber struct {
	ID                  int          `json:"id"`
	Number              Phone        `json:"number"`
	VoiceLanguageLocale string       `json:"voiceLanguageLocale"`
	AlertSource         *AlertSource `json:"alertSource"`
}

// IncidentStatuses defines incident statuses
var IncidentStatuses = struct {
	New      string
	Pending  string
	Accepted string
	Resolved string
}{
	New:      "NEW",
	Pending:  "PENDING",
	Accepted: "ACCEPTED",
	Resolved: "RESOLVED",
}

// IncidentPriorities defines incident priorities
var IncidentPriorities = struct {
	High string
	Low  string
}{
	High: "HIGH",
	Low:  "LOW",
}

// IncidentResponderTypes defines incident responder types
var IncidentResponderTypes = struct {
	User        string
	AlertSource string
}{
	User:        "USER",
	AlertSource: "SOURCE",
}

// IncidentResponder definition
type IncidentResponder struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Group    string `json:"group"`
	Disabled bool   `json:"disabled"`
}

// IncidentResponderGroups defines incident responder groups
var IncidentResponderGroups = struct {
	Suggested        string
	User             string
	EscalationPolicy string
	OnCallSchedule   string
}{
	Suggested:        "SUGGESTED",
	User:             "USER",
	EscalationPolicy: "ESCALATION_POLICY",
	OnCallSchedule:   "ON_CALL_SCHEDULE",
}

// IncidentLogEntry definition
type IncidentLogEntry struct {
	ID           int64  `json:"id"`
	Timestamp    string `json:"timestamp"` // Date time string in ISO format
	LogEntryType string `json:"logEntryType"`
	Text         string `json:"text"`
	IncidentID   int64  `json:"incidentId"`
}

// IncidentLogEntryTypes defines incident log entry types
var IncidentLogEntryTypes = struct {
	AlertReceivedLogEntry            string
	AlertSourceResponseLogEntry      string
	EmailReceivedLogEntry            string
	IncidentAssignedBySystemLogEntry string
	IncidentAssignedByUserLogEntry   string
	IncidentCreatedByUserLogEntry    string
	NotificationLogEntry             string
	UserResponseLogEntry             string
}{
	AlertReceivedLogEntry:            "AlertReceivedLogEntry",
	AlertSourceResponseLogEntry:      "AlertSourceResponseLogEntry",
	EmailReceivedLogEntry:            "EmailReceivedLogEntry",
	IncidentAssignedBySystemLogEntry: "IncidentAssignedBySystemLogEntry",
	IncidentAssignedByUserLogEntry:   "IncidentAssignedByUserLogEntry",
	IncidentCreatedByUserLogEntry:    "IncidentCreatedByUserLogEntry",
	NotificationLogEntry:             "NotificationLogEntry",
	UserResponseLogEntry:             "UserResponseLogEntry",
}

// IncidentAction definition
type IncidentAction struct {
	Name        string                 `json:"name"`
	WebhookID   string                 `json:"webhookId"`
	ExtensionID string                 `json:"extensionId"`
	IconURL     string                 `json:"iconUrl"`
	History     []IncidentActionResult `json:"history"`
}

// IncidentActionResult definition
type IncidentActionResult struct {
	ID          string `json:"id"`
	IncidentID  int64  `json:"incidentId"`
	WebhookID   string `json:"webhookId"`
	ExtensionID string `json:"extensionId"`
	Actor       User   `json:"actor"`
	Success     bool   `json:"success"`
}

// GetIncidentInput represents the input of a GetIncident operation.
type GetIncidentInput struct {
	_          struct{}
	IncidentID *int64
}

// GetIncidentOutput represents the output of a GetIncident operation.
type GetIncidentOutput struct {
	_        struct{}
	Incident *Incident
}

// GetIncident gets the incident with specified id. https://api.ilert.com/api-docs/#tag/Incidents/paths/~1incidents~1{id}/get
func (c *Client) GetIncident(input *GetIncidentInput) (*GetIncidentOutput, error) {
	if input == nil {
		return nil, errors.New("input is required")
	}
	if input.IncidentID == nil {
		return nil, errors.New("Incident id is required")
	}

	resp, err := c.httpClient.R().Get(fmt.Sprintf("%s/%d", apiRoutes.incidents, *input.IncidentID))
	if err != nil {
		return nil, err
	}
	if apiErr := getGenericAPIError(resp, 200); apiErr != nil {
		return nil, apiErr
	}

	incident := &Incident{}
	err = json.Unmarshal(resp.Body(), incident)
	if err != nil {
		return nil, err
	}

	return &GetIncidentOutput{Incident: incident}, nil
}

// GetIncidentsInput represents the input of a GetIncidents operation.
type GetIncidentsInput struct {
	_ struct{}
	// an integer specifying the starting point (beginning with 0) when paging through a list of entities
	StartIndex *int

	// the maximum number of results when paging through a list of entities.
	// Default: 50
	MaxResults *int

	// state of the incident
	States []*string

	// alert source IDs of the incident's alert source
	AlertSources []*int64

	// user IDs of the user that the incident is assigned to
	AssignedToUserIDs []*int64

	// usernames of the user that the incident is assigned to
	AssignedToUserNames []*string

	// Date time string in ISO format
	From *string

	// Date time string in ISO format
	Until *string
}

// GetIncidentsOutput represents the output of a GetIncidents operation.
type GetIncidentsOutput struct {
	_         struct{}
	Incidents []*Incident
}

// GetIncidents lists alert sources. https://api.ilert.com/api-docs/#tag/Incidents/paths/~1incidents/get
func (c *Client) GetIncidents(input *GetIncidentsInput) (*GetIncidentsOutput, error) {
	if input == nil {
		input = &GetIncidentsInput{}
	}

	q := url.Values{}
	if input.StartIndex != nil {
		q.Add("start-index", strconv.Itoa(*input.StartIndex))
	}
	if input.MaxResults != nil {
		q.Add("max-results", strconv.Itoa(*input.MaxResults))
	}
	if input.From != nil {
		q.Add("from", *input.From)
	}
	if input.Until != nil {
		q.Add("until", *input.From)
	}

	for _, state := range input.States {
		q.Add("state", *state)
	}

	for _, alertSourceID := range input.AlertSources {
		q.Add("alert-source", strconv.FormatInt(*alertSourceID, 10))
	}

	for _, userID := range input.AssignedToUserIDs {
		q.Add("assigned-to", strconv.FormatInt(*userID, 10))
	}

	for _, username := range input.AssignedToUserNames {
		q.Add("assigned-to", *username)
	}

	resp, err := c.httpClient.R().Get(fmt.Sprintf("%s?%s", apiRoutes.incidents, q.Encode()))
	if err != nil {
		return nil, err
	}
	if apiErr := getGenericAPIError(resp, 200); apiErr != nil {
		return nil, apiErr
	}

	incidents := make([]*Incident, 0)
	err = json.Unmarshal(resp.Body(), &incidents)
	if err != nil {
		return nil, err
	}

	return &GetIncidentsOutput{Incidents: incidents}, nil
}

// GetIncidentsCountInput represents the input of a GetIncidentsCount operation.
type GetIncidentsCountInput struct {
	_ struct{}

	// state of the incident
	States []*string

	// alert source IDs of the incident's alert source
	AlertSources []*int64

	// user IDs of the user that the incident is assigned to
	AssignedToUserIDs []*int64

	// usernames of the user that the incident is assigned to
	AssignedToUserNames []*string

	// Date time string in ISO format
	From *string

	// Date time string in ISO format
	Until *string
}

// GetIncidentsCountOutput represents the output of a GetIncidentsCount operation.
type GetIncidentsCountOutput struct {
	_     struct{}
	Count int
}

// GetIncidentsCount gets list uptime monitors. https://api.ilert.com/api-docs/#tag/Incidents/paths/~1incidents~1count/get
func (c *Client) GetIncidentsCount(input *GetIncidentsCountInput) (*GetIncidentsCountOutput, error) {
	if input == nil {
		input = &GetIncidentsCountInput{}
	}

	q := url.Values{}
	if input.From != nil {
		q.Add("from", *input.From)
	}
	if input.Until != nil {
		q.Add("until", *input.From)
	}

	for _, state := range input.States {
		q.Add("state", *state)
	}

	for _, alertSourceID := range input.AlertSources {
		q.Add("alert-source", strconv.FormatInt(*alertSourceID, 10))
	}

	for _, userID := range input.AssignedToUserIDs {
		q.Add("assigned-to", strconv.FormatInt(*userID, 10))
	}

	for _, username := range input.AssignedToUserNames {
		q.Add("assigned-to", *username)
	}

	resp, err := c.httpClient.R().Get(fmt.Sprintf("%s/count?%s", apiRoutes.incidents, q.Encode()))
	if err != nil {
		return nil, err
	}
	if apiErr := getGenericAPIError(resp, 200); apiErr != nil {
		return nil, apiErr
	}

	body := &GenericCountResponse{}
	err = json.Unmarshal(resp.Body(), body)
	if err != nil {
		return nil, err
	}

	return &GetIncidentsCountOutput{Count: body.Count}, nil
}

// GetIncidentResponderInput represents the input of a GetIncidentResponder operation.
type GetIncidentResponderInput struct {
	_          struct{}
	IncidentID *int64
	Language   *string
}

// GetIncidentResponderOutput represents the output of a GetIncidentResponder operation.
type GetIncidentResponderOutput struct {
	_          struct{}
	Responders []*IncidentResponder
}

// GetIncidentResponder gets the alert source with specified id. https://api.ilert.com/api-docs/#tag/Incidents/paths/~1incidents~1{id}~1responder/get
func (c *Client) GetIncidentResponder(input *GetIncidentResponderInput) (*GetIncidentResponderOutput, error) {
	if input == nil {
		return nil, errors.New("input is required")
	}
	if input.IncidentID == nil {
		return nil, errors.New("Incident id is required")
	}

	q := url.Values{}
	if input.Language != nil {
		if *input.Language == "en" {
			q.Add("lng", "en")
		} else if *input.Language == "de" {
			q.Add("lng", "de")
		}
	}

	resp, err := c.httpClient.R().Get(fmt.Sprintf("%s/%d/responder", apiRoutes.incidents, *input.IncidentID))
	if err != nil {
		return nil, err
	}
	if apiErr := getGenericAPIError(resp, 200); apiErr != nil {
		return nil, apiErr
	}

	incidentResponders := make([]*IncidentResponder, 0)
	err = json.Unmarshal(resp.Body(), &incidentResponders)
	if err != nil {
		return nil, err
	}

	return &GetIncidentResponderOutput{Responders: incidentResponders}, nil
}

// AssignIncidentInput represents the input of a AssignIncident operation.
type AssignIncidentInput struct {
	_                  struct{}
	IncidentID         *int64
	UserID             *int64
	Username           *string
	EscalationPolicyID *int64
	ScheduleID         *int64
}

// AssignIncidentOutput represents the output of a AssignIncident operation.
type AssignIncidentOutput struct {
	_        struct{}
	Incident *Incident
}

// AssignIncident gets the alert source with specified id. https://api.ilert.com/api-docs/#tag/Incidents/paths/~1incidents~1{id}~1assign/put
func (c *Client) AssignIncident(input *AssignIncidentInput) (*AssignIncidentOutput, error) {
	if input == nil {
		return nil, errors.New("input is required")
	}
	if input.IncidentID == nil {
		return nil, errors.New("Incident id is required")
	}

	if input.UserID == nil && input.Username == nil && input.EscalationPolicyID == nil && input.ScheduleID == nil {
		return nil, errors.New("one of assignments is required")
	}

	q := url.Values{}
	if input.UserID != nil {
		q.Add("user-id", strconv.FormatInt(*input.UserID, 10))
	}
	if input.Username != nil {
		q.Add("user-id", *input.Username)
	}
	if input.EscalationPolicyID != nil {
		q.Add("policy-id", strconv.FormatInt(*input.EscalationPolicyID, 10))
	}
	if input.ScheduleID != nil {
		q.Add("schedule-id", strconv.FormatInt(*input.ScheduleID, 10))
	}

	resp, err := c.httpClient.R().Put(fmt.Sprintf("%s/%d/assign?%s", apiRoutes.incidents, *input.IncidentID, q.Encode()))
	if err != nil {
		return nil, err
	}
	if apiErr := getGenericAPIError(resp, 200); apiErr != nil {
		return nil, apiErr
	}

	incident := &Incident{}
	err = json.Unmarshal(resp.Body(), incident)
	if err != nil {
		return nil, err
	}

	return &AssignIncidentOutput{Incident: incident}, nil
}

// AcceptIncidentInput represents the input of a AcceptIncident operation.
type AcceptIncidentInput struct {
	_          struct{}
	IncidentID *int64
}

// AcceptIncidentOutput represents the output of a AcceptIncident operation.
type AcceptIncidentOutput struct {
	_        struct{}
	Incident *Incident
}

// AcceptIncident gets the alert source with specified id. https://api.ilert.com/api-docs/#tag/Incidents/paths/~1incidents~1{id}~1accept/put
func (c *Client) AcceptIncident(input *AcceptIncidentInput) (*AcceptIncidentOutput, error) {
	if input == nil {
		return nil, errors.New("input is required")
	}
	if input.IncidentID == nil {
		return nil, errors.New("Incident id is required")
	}

	resp, err := c.httpClient.R().Put(fmt.Sprintf("%s/%d/accept", apiRoutes.incidents, *input.IncidentID))
	if err != nil {
		return nil, err
	}
	if apiErr := getGenericAPIError(resp, 200); apiErr != nil {
		return nil, apiErr
	}

	incident := &Incident{}
	err = json.Unmarshal(resp.Body(), incident)
	if err != nil {
		return nil, err
	}

	return &AcceptIncidentOutput{Incident: incident}, nil
}

// ResolveIncidentInput represents the input of a ResolveIncident operation.
type ResolveIncidentInput struct {
	_          struct{}
	IncidentID *int64
}

// ResolveIncidentOutput represents the output of a ResolveIncident operation.
type ResolveIncidentOutput struct {
	_        struct{}
	Incident *Incident
}

// ResolveIncident gets the alert source with specified id. https://api.ilert.com/api-docs/#tag/Incidents/paths/~1incidents~1{id}~1resolve/put
func (c *Client) ResolveIncident(input *ResolveIncidentInput) (*ResolveIncidentOutput, error) {
	if input == nil {
		return nil, errors.New("input is required")
	}
	if input.IncidentID == nil {
		return nil, errors.New("Incident id is required")
	}

	resp, err := c.httpClient.R().Put(fmt.Sprintf("%s/%d/resolve", apiRoutes.incidents, *input.IncidentID))
	if err != nil {
		return nil, err
	}
	if apiErr := getGenericAPIError(resp, 200); apiErr != nil {
		return nil, apiErr
	}

	incident := &Incident{}
	err = json.Unmarshal(resp.Body(), incident)
	if err != nil {
		return nil, err
	}

	return &ResolveIncidentOutput{Incident: incident}, nil
}

// GetIncidentLogEntriesInput represents the input of a GetIncidentLogEntries operation.
type GetIncidentLogEntriesInput struct {
	_          struct{}
	IncidentID *int64
	Language   *string
}

// GetIncidentLogEntriesOutput represents the output of a GetIncidentLogEntries operation.
type GetIncidentLogEntriesOutput struct {
	_          struct{}
	LogEntries []*IncidentLogEntry
}

// GetIncidentLogEntries gets log entries for the specified incident. https://api.ilert.com/api-docs/#tag/Incidents/paths/~1incidents~1{id}~1log-entries/get
func (c *Client) GetIncidentLogEntries(input *GetIncidentLogEntriesInput) (*GetIncidentLogEntriesOutput, error) {
	if input == nil {
		return nil, errors.New("input is required")
	}
	if input.IncidentID == nil {
		return nil, errors.New("Incident id is required")
	}

	q := url.Values{}
	if input.Language != nil {
		if *input.Language == "en" {
			q.Add("lng", "en")
		} else if *input.Language == "de" {
			q.Add("lng", "de")
		}
	}

	resp, err := c.httpClient.R().Get(fmt.Sprintf("%s/%d/log-entries", apiRoutes.incidents, *input.IncidentID))
	if err != nil {
		return nil, err
	}
	if apiErr := getGenericAPIError(resp, 200); apiErr != nil {
		return nil, apiErr
	}

	incidentLogEntries := make([]*IncidentLogEntry, 0)
	err = json.Unmarshal(resp.Body(), &incidentLogEntries)
	if err != nil {
		return nil, err
	}

	return &GetIncidentLogEntriesOutput{LogEntries: incidentLogEntries}, nil
}

// TODO https://api.ilert.com/api-docs/#tag/Incidents/paths/~1incidents~1{id}~1notifications/get

// GetIncidentActionsInput represents the input of a GetIncidentsAction operation.
type GetIncidentActionsInput struct {
	_          struct{}
	IncidentID *int64
}

// GetIncidentActionsOutput represents the output of a GetIncidentsAction operation.
type GetIncidentActionsOutput struct {
	_       struct{}
	Actions []*IncidentAction
}

// GetIncidentActions gets the incident with specified id. https://api.ilert.com/api-docs/#tag/Incident-Actions/paths/~1incidents~1{id}~1actions/get
func (c *Client) GetIncidentActions(input *GetIncidentActionsInput) (*GetIncidentActionsOutput, error) {
	if input == nil {
		return nil, errors.New("input is required")
	}
	if input.IncidentID == nil {
		return nil, errors.New("Incident id is required")
	}

	resp, err := c.httpClient.R().Get(fmt.Sprintf("%s/%d/actions", apiRoutes.incidents, *input.IncidentID))
	if err != nil {
		return nil, err
	}
	if apiErr := getGenericAPIError(resp, 200); apiErr != nil {
		return nil, apiErr
	}

	incidentActions := make([]*IncidentAction, 0)
	err = json.Unmarshal(resp.Body(), &incidentActions)
	if err != nil {
		return nil, err
	}

	return &GetIncidentActionsOutput{Actions: incidentActions}, nil
}

// InvokeIncidentActionInput represents the input of a InvokeIncidentAction operation.
type InvokeIncidentActionInput struct {
	_          struct{}
	IncidentID *int64
	Action     *IncidentAction
}

// InvokeIncidentActionOutput represents the output of a InvokeIncidentAction operation.
type InvokeIncidentActionOutput struct {
	_      struct{}
	Action *IncidentAction
}

// InvokeIncidentAction creates a new alert source. https://api.ilert.com/api-docs/#tag/Incident-Actions/paths/~1incidents~1{id}~1actions/post
func (c *Client) InvokeIncidentAction(input *InvokeIncidentActionInput) (*InvokeIncidentActionOutput, error) {
	if input == nil {
		return nil, errors.New("input is required")
	}
	if input.IncidentID == nil {
		return nil, errors.New("Incident id is required")
	}
	if input.Action == nil {
		return nil, errors.New("action input is required")
	}

	resp, err := c.httpClient.R().SetBody(input.Action).Post(fmt.Sprintf("%s/%d/actions", apiRoutes.incidents, *input.IncidentID))
	if err != nil {
		return nil, err
	}
	if apiErr := getGenericAPIError(resp, 201); apiErr != nil {
		return nil, apiErr
	}

	incidentAction := &IncidentAction{}
	err = json.Unmarshal(resp.Body(), incidentAction)
	if err != nil {
		return nil, err
	}

	return &InvokeIncidentActionOutput{Action: incidentAction}, nil
}
