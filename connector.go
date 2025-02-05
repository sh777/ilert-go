package ilert

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Connector definition
type Connector struct {
	ID        string      `json:"id,omitempty"`
	Name      string      `json:"name"`
	Type      string      `json:"type"`
	CreatedAt string      `json:"createdAt,omitempty"` // date time string in ISO 8601
	UpdatedAt string      `json:"updatedAt,omitempty"` // date time string in ISO 8601
	Params    interface{} `json:"params"`
}

// ConnectorOutput definition
type ConnectorOutput struct {
	ID        string                `json:"id"`
	Name      string                `json:"name"`
	Type      string                `json:"type"`
	CreatedAt string                `json:"createdAt"` // date time string in ISO 8601
	UpdatedAt string                `json:"updatedAt"` // date time string in ISO 8601
	Params    ConnectorOutputParams `json:"params"`
}

// ConnectorOutputParams definition
type ConnectorOutputParams struct {
	APIKey        string `json:"apiKey,omitempty"`        // Datadog or Zendesk or Github or Serverless or Autotask api key
	Authorization string `json:"authorization,omitempty"` // Serverless
	URL           string `json:"url,omitempty"`           // Jira or Microsoft Teams or Zendesk or Discord or Autotask server url
	Email         string `json:"email,omitempty"`         // Jira or ServiceNow or Zendesk username or email
	Username      string `json:"username,omitempty"`      // TOPdesk or ServiceNow or Autotask username
	Password      string `json:"password,omitempty"`      // Jira or ServiceNow or Autotask user password or api token
}

// ConnectorParamsDatadog definition
type ConnectorParamsDatadog struct {
	APIKey string `json:"apiKey"`
}

// ConnectorParamsJira definition
type ConnectorParamsJira struct {
	URL      string `json:"url"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ConnectorParamsMicrosoftTeams definition
type ConnectorParamsMicrosoftTeams struct {
	URL string `json:"url"`
}

// ConnectorParamsServiceNow definition
type ConnectorParamsServiceNow struct {
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// ConnectorParamsSlack definition
type ConnectorParamsSlack struct{}

// ConnectorParamsZendesk definition
type ConnectorParamsZendesk struct {
	URL    string `json:"url"`
	Email  string `json:"email"`
	APIKey string `json:"apiKey"`
}

// ConnectorParamsDiscord definition
type ConnectorParamsDiscord struct {
	URL string `json:"url"`
}

// ConnectorParamsGithub definition
type ConnectorParamsGithub struct {
	APIKey string `json:"apiKey"`
}

// ConnectorParamsTopdesk definition
type ConnectorParamsTopdesk struct {
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// ConnectorParamsAWSLambda definition
type ConnectorParamsAWSLambda struct {
	Authorization string `json:"authorization,omitempty"`
}

// ConnectorParamsAzureFunction definition
type ConnectorParamsAzureFunction struct {
	Authorization string `json:"authorization,omitempty"`
}

// ConnectorParamsGoogleFunction definition
type ConnectorParamsGoogleFunction struct {
	Authorization string `json:"authorization,omitempty"`
}

// ConnectorParamsSysdig definition
type ConnectorParamsSysdig struct {
	APIKey string `json:"apiKey"`
}

// ConnectorParamsAutotask definition
type ConnectorParamsAutotask struct {
	URL      string `json:"url"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ConnectorParamsMattermost definition
type ConnectorParamsMattermost struct {
	URL string `json:"url"`
}

// ConnectorParamsZammad definition
type ConnectorParamsZammad struct {
	URL    string `json:"url"`
	APIKey string `json:"apiKey"`
}

// ConnectorParamsStatusPageIO definition
type ConnectorParamsStatusPageIO struct {
	APIKey string `json:"apiKey"`
}

// ConnectorTypes defines connector types
var ConnectorTypes = struct {
	AWSLambda             string
	AzureFAAS             string
	Datadog               string
	Discord               string
	Email                 string
	Github                string
	GoogleFAAS            string
	Jira                  string
	MicrosoftTeams        string
	ServiceNow            string
	Slack                 string
	Sysdig                string
	Topdesk               string
	Webhook               string
	Zapier                string
	Zendesk               string
	MicrosoftTeamsChat    string
	MicrosoftTeamsMeeting string
	Autotask              string
	Mattermost            string
	Zammad                string
	ZoomChat              string
	ZoomMeeting           string
	StatusPageIO          string
	Webex                 string
}{
	AWSLambda:             "aws_lambda",
	AzureFAAS:             "azure_faas",
	Datadog:               "datadog",
	Discord:               "discord",
	Email:                 "email",
	Github:                "github",
	GoogleFAAS:            "google_faas",
	Jira:                  "jira",
	MicrosoftTeams:        "microsoft_teams",
	ServiceNow:            "servicenow",
	Slack:                 "slack",
	Sysdig:                "sysdig",
	Topdesk:               "topdesk",
	Webhook:               "webhook",
	Zapier:                "zapier",
	Zendesk:               "zendesk",
	MicrosoftTeamsChat:    "microsoft_teams_chat",
	MicrosoftTeamsMeeting: "microsoft_teams_meeting",
	Autotask:              "autotask",
	Mattermost:            "mattermost",
	Zammad:                "zammad",
	ZoomChat:              "zoom_chat",
	ZoomMeeting:           "zoom_meeting",
	StatusPageIO:          "status_page_io",
	Webex:                 "webex",
}

// ConnectorTypesAll defines connector all types list
var ConnectorTypesAll = []string{
	ConnectorTypes.AWSLambda,
	ConnectorTypes.AzureFAAS,
	ConnectorTypes.Datadog,
	ConnectorTypes.Discord,
	ConnectorTypes.Email,
	ConnectorTypes.Github,
	ConnectorTypes.GoogleFAAS,
	ConnectorTypes.Jira,
	ConnectorTypes.MicrosoftTeams,
	ConnectorTypes.ServiceNow,
	ConnectorTypes.Slack,
	ConnectorTypes.Sysdig,
	ConnectorTypes.Topdesk,
	ConnectorTypes.Webhook,
	ConnectorTypes.Zapier,
	ConnectorTypes.Zendesk,
	ConnectorTypes.MicrosoftTeamsChat,
	ConnectorTypes.MicrosoftTeamsMeeting,
	ConnectorTypes.Autotask,
	ConnectorTypes.Mattermost,
	ConnectorTypes.Zammad,
	ConnectorTypes.ZoomChat,
	ConnectorTypes.ZoomMeeting,
	ConnectorTypes.StatusPageIO,
	ConnectorTypes.Webex,
}

// CreateConnectorInput represents the input of a CreateConnector operation.
type CreateConnectorInput struct {
	_         struct{}
	Connector *Connector
}

// CreateConnectorOutput represents the output of a CreateConnector operation.
type CreateConnectorOutput struct {
	_         struct{}
	Connector *ConnectorOutput
}

// CreateConnector creates a new connector. https://api.ilert.com/api-docs/#tag/Connectors/paths/~1connectors/post
func (c *Client) CreateConnector(input *CreateConnectorInput) (*CreateConnectorOutput, error) {
	if input == nil {
		return nil, errors.New("input is required")
	}
	if input.Connector == nil {
		return nil, errors.New("Connector input is required")
	}
	resp, err := c.httpClient.R().SetBody(input.Connector).Post(apiRoutes.connectors)
	if err != nil {
		return nil, err
	}
	if apiErr := getGenericAPIError(resp, 201); apiErr != nil {
		return nil, apiErr
	}

	connector := &ConnectorOutput{}
	err = json.Unmarshal(resp.Body(), connector)
	if err != nil {
		return nil, err
	}

	return &CreateConnectorOutput{Connector: connector}, nil
}

// GetConnectorInput represents the input of a GetConnector operation.
type GetConnectorInput struct {
	_           struct{}
	ConnectorID *string
}

// GetConnectorOutput represents the output of a GetConnector operation.
type GetConnectorOutput struct {
	_         struct{}
	Connector *ConnectorOutput
}

// GetConnector gets the connector with specified id. https://api.ilert.com/api-docs/#tag/Connectors/paths/~1connectors~1{id}/get
func (c *Client) GetConnector(input *GetConnectorInput) (*GetConnectorOutput, error) {
	if input == nil {
		return nil, errors.New("input is required")
	}
	if input.ConnectorID == nil {
		return nil, errors.New("Connector id is required")
	}

	resp, err := c.httpClient.R().Get(fmt.Sprintf("%s/%s", apiRoutes.connectors, *input.ConnectorID))
	if err != nil {
		return nil, err
	}
	if apiErr := getGenericAPIError(resp, 200); apiErr != nil {
		return nil, apiErr
	}

	connector := &ConnectorOutput{}
	err = json.Unmarshal(resp.Body(), connector)
	if err != nil {
		return nil, err
	}

	return &GetConnectorOutput{Connector: connector}, nil
}

// GetConnectorsInput represents the input of a GetConnectors operation.
type GetConnectorsInput struct {
	_ struct{}
}

// GetConnectorsOutput represents the output of a GetConnectors operation.
type GetConnectorsOutput struct {
	_          struct{}
	Connectors []*ConnectorOutput
}

// GetConnectors lists connectors. https://api.ilert.com/api-docs/#tag/Connectors/paths/~1connectors/get
func (c *Client) GetConnectors(input *GetConnectorsInput) (*GetConnectorsOutput, error) {
	resp, err := c.httpClient.R().Get(apiRoutes.connectors)
	if err != nil {
		return nil, err
	}
	if apiErr := getGenericAPIError(resp, 200); apiErr != nil {
		return nil, apiErr
	}

	connectors := make([]*ConnectorOutput, 0)
	err = json.Unmarshal(resp.Body(), &connectors)
	if err != nil {
		return nil, err
	}

	return &GetConnectorsOutput{Connectors: connectors}, nil
}

// UpdateConnectorInput represents the input of a UpdateConnector operation.
type UpdateConnectorInput struct {
	_           struct{}
	ConnectorID *string
	Connector   *Connector
}

// UpdateConnectorOutput represents the output of a UpdateConnector operation.
type UpdateConnectorOutput struct {
	_         struct{}
	Connector *ConnectorOutput
}

// UpdateConnector updates an existing connector. https://api.ilert.com/api-docs/#tag/Connectors/paths/~1connectors~1{id}/put
func (c *Client) UpdateConnector(input *UpdateConnectorInput) (*UpdateConnectorOutput, error) {
	if input == nil {
		return nil, errors.New("input is required")
	}
	if input.Connector == nil {
		return nil, errors.New("Connector input is required")
	}
	if input.ConnectorID == nil {
		return nil, errors.New("Connector id is required")
	}

	resp, err := c.httpClient.R().SetBody(input.Connector).Put(fmt.Sprintf("%s/%s", apiRoutes.connectors, *input.ConnectorID))
	if err != nil {
		return nil, err
	}
	if apiErr := getGenericAPIError(resp, 200); apiErr != nil {
		return nil, apiErr
	}

	connector := &ConnectorOutput{}
	err = json.Unmarshal(resp.Body(), connector)
	if err != nil {
		return nil, err
	}

	return &UpdateConnectorOutput{Connector: connector}, nil
}

// DeleteConnectorInput represents the input of a DeleteConnector operation.
type DeleteConnectorInput struct {
	_           struct{}
	ConnectorID *string
}

// DeleteConnectorOutput represents the output of a DeleteConnector operation.
type DeleteConnectorOutput struct {
	_ struct{}
}

// DeleteConnector deletes the specified alert source. https://api.ilert.com/api-docs/#tag/Connectors/paths/~1connectors~1{id}/delete
func (c *Client) DeleteConnector(input *DeleteConnectorInput) (*DeleteConnectorOutput, error) {
	if input == nil {
		return nil, errors.New("input is required")
	}
	if input.ConnectorID == nil {
		return nil, errors.New("Connector id is required")
	}

	resp, err := c.httpClient.R().Delete(fmt.Sprintf("%s/%s", apiRoutes.connectors, *input.ConnectorID))
	if err != nil {
		return nil, err
	}
	if apiErr := getGenericAPIError(resp, 204); apiErr != nil {
		return nil, apiErr
	}

	return &DeleteConnectorOutput{}, nil
}
