package ilert

import (
	"encoding/json"
	"errors"
	"fmt"
)

// EscalationPolicy definition https://api.ilert.com/api-docs/#!/Escalation_Policies
type EscalationPolicy struct {
	ID              int64            `json:"id,omitempty"`
	Name            string           `json:"name"`
	EscalationRules []EscalationRule `json:"escalationRules"`
	Repeating       bool             `json:"repeating,omitempty"`
	Frequency       int              `json:"frequency,omitempty"`
	Teams           []TeamShort      `json:"teams,omitempty"`
}

// EscalationRule definition
type EscalationRule struct {
	User              *User     `json:"user,omitempty"`
	Schedule          *Schedule `json:"schedule,omitempty"`
	EscalationTimeout int       `json:"escalationTimeout"`
}

// CreateEscalationPolicyInput represents the input of a CreateEscalationPolicy operation.
type CreateEscalationPolicyInput struct {
	_                struct{}
	EscalationPolicy *EscalationPolicy
}

// CreateEscalationPolicyOutput represents the output of a CreateEscalationPolicy operation.
type CreateEscalationPolicyOutput struct {
	_                struct{}
	EscalationPolicy *EscalationPolicy
}

// CreateEscalationPolicy creates a new escalation policy. https://api.ilert.com/api-docs/#tag/Escalation-Policies/paths/~1escalation-policies/post
func (c *Client) CreateEscalationPolicy(input *CreateEscalationPolicyInput) (*CreateEscalationPolicyOutput, error) {
	if input == nil {
		return nil, errors.New("input is required")
	}
	if input.EscalationPolicy == nil {
		return nil, errors.New("escalation policy input is required")
	}
	resp, err := c.httpClient.R().SetBody(input.EscalationPolicy).Post(apiRoutes.escalationPolicies)
	if err != nil {
		return nil, err
	}
	if apiErr := getGenericAPIError(resp, 201); apiErr != nil {
		return nil, apiErr
	}

	escalationPolicy := &EscalationPolicy{}
	err = json.Unmarshal(resp.Body(), escalationPolicy)
	if err != nil {
		return nil, err
	}

	return &CreateEscalationPolicyOutput{EscalationPolicy: escalationPolicy}, nil
}

// GetEscalationPolicyInput represents the input of a GetEscalationPolicy operation.
type GetEscalationPolicyInput struct {
	_                  struct{}
	EscalationPolicyID *int64
}

// GetEscalationPolicyOutput represents the output of a GetEscalationPolicy operation.
type GetEscalationPolicyOutput struct {
	_                struct{}
	EscalationPolicy *EscalationPolicy
}

// GetEscalationPolicy gets the escalation policy with specified id. https://api.ilert.com/api-docs/#tag/Escalation-Policies/paths/~1escalation-policies~1{id}/get
func (c *Client) GetEscalationPolicy(input *GetEscalationPolicyInput) (*GetEscalationPolicyOutput, error) {
	if input == nil {
		return nil, errors.New("input is required")
	}
	if input.EscalationPolicyID == nil {
		return nil, errors.New("EscalationPolicy id is required")
	}

	resp, err := c.httpClient.R().Get(fmt.Sprintf("%s/%d", apiRoutes.escalationPolicies, *input.EscalationPolicyID))
	if err != nil {
		return nil, err
	}
	if apiErr := getGenericAPIError(resp, 200); apiErr != nil {
		return nil, apiErr
	}

	escalationPolicy := &EscalationPolicy{}
	err = json.Unmarshal(resp.Body(), escalationPolicy)
	if err != nil {
		return nil, err
	}

	return &GetEscalationPolicyOutput{EscalationPolicy: escalationPolicy}, nil
}

// GetEscalationPoliciesInput represents the input of a GetEscalationPolicies operation.
type GetEscalationPoliciesInput struct {
	_ struct{}
}

// GetEscalationPoliciesOutput represents the output of a GetEscalationPolicies operation.
type GetEscalationPoliciesOutput struct {
	_                  struct{}
	EscalationPolicies []*EscalationPolicy
}

// GetEscalationPolicies lists escalation policies. https://api.ilert.com/api-docs/#tag/Escalation-Policies/paths/~1escalation-policies/get
func (c *Client) GetEscalationPolicies(input *GetEscalationPoliciesInput) (*GetEscalationPoliciesOutput, error) {
	resp, err := c.httpClient.R().Get(apiRoutes.escalationPolicies)
	if err != nil {
		return nil, err
	}
	if apiErr := getGenericAPIError(resp, 200); apiErr != nil {
		return nil, apiErr
	}

	escalationPolicies := make([]*EscalationPolicy, 0)
	err = json.Unmarshal(resp.Body(), &escalationPolicies)
	if err != nil {
		return nil, err
	}

	return &GetEscalationPoliciesOutput{EscalationPolicies: escalationPolicies}, nil
}

// UpdateEscalationPolicyInput represents the input of a UpdateEscalationPolicy operation.
type UpdateEscalationPolicyInput struct {
	_                  struct{}
	EscalationPolicyID *int64
	EscalationPolicy   *EscalationPolicy
}

// UpdateEscalationPolicyOutput represents the output of a UpdateEscalationPolicy operation.
type UpdateEscalationPolicyOutput struct {
	_                struct{}
	EscalationPolicy *EscalationPolicy
}

// UpdateEscalationPolicy updates an existing escalation policy. https://api.ilert.com/api-docs/#tag/Escalation-Policies/paths/~1escalation-policies~1{id}/put
func (c *Client) UpdateEscalationPolicy(input *UpdateEscalationPolicyInput) (*UpdateEscalationPolicyOutput, error) {
	if input == nil {
		return nil, errors.New("input is required")
	}
	if input.EscalationPolicy == nil {
		return nil, errors.New("EscalationPolicy input is required")
	}
	if input.EscalationPolicyID == nil {
		return nil, errors.New("escalation policy id is required")
	}

	resp, err := c.httpClient.R().SetBody(input.EscalationPolicy).Put(fmt.Sprintf("%s/%d", apiRoutes.escalationPolicies, *input.EscalationPolicyID))
	if err != nil {
		return nil, err
	}
	if apiErr := getGenericAPIError(resp, 200); apiErr != nil {
		return nil, apiErr
	}

	escalationPolicy := &EscalationPolicy{}
	err = json.Unmarshal(resp.Body(), escalationPolicy)
	if err != nil {
		return nil, err
	}

	return &UpdateEscalationPolicyOutput{EscalationPolicy: escalationPolicy}, nil
}

// DeleteEscalationPolicyInput represents the input of a DeleteEscalationPolicy operation.
type DeleteEscalationPolicyInput struct {
	_                  struct{}
	EscalationPolicyID *int64
}

// DeleteEscalationPolicyOutput represents the output of a DeleteEscalationPolicy operation.
type DeleteEscalationPolicyOutput struct {
	_ struct{}
}

// DeleteEscalationPolicy deletes the specified escalation policy. https://api.ilert.com/api-docs/#tag/Escalation-Policies/paths/~1escalation-policies~1{id}/delete
func (c *Client) DeleteEscalationPolicy(input *DeleteEscalationPolicyInput) (*DeleteEscalationPolicyOutput, error) {
	if input == nil {
		return nil, errors.New("input is required")
	}
	if input.EscalationPolicyID == nil {
		return nil, errors.New("EscalationPolicy id is required")
	}

	resp, err := c.httpClient.R().Delete(fmt.Sprintf("%s/%d", apiRoutes.escalationPolicies, *input.EscalationPolicyID))
	if err != nil {
		return nil, err
	}
	if apiErr := getGenericAPIError(resp, 204); apiErr != nil {
		return nil, apiErr
	}

	output := &DeleteEscalationPolicyOutput{}
	return output, nil
}
