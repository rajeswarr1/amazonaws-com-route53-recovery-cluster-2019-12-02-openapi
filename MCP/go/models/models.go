package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// RoutingControl represents the RoutingControl schema from the OpenAPI specification
type RoutingControl struct {
	Controlpanelarn interface{} `json:"ControlPanelArn,omitempty"`
	Controlpanelname interface{} `json:"ControlPanelName,omitempty"`
	Routingcontrolarn interface{} `json:"RoutingControlArn,omitempty"`
	Routingcontrolname interface{} `json:"RoutingControlName,omitempty"`
	Routingcontrolstate interface{} `json:"RoutingControlState,omitempty"`
}

// ListRoutingControlsResponse represents the ListRoutingControlsResponse schema from the OpenAPI specification
type ListRoutingControlsResponse struct {
	Routingcontrols interface{} `json:"RoutingControls"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// GetRoutingControlStateRequest represents the GetRoutingControlStateRequest schema from the OpenAPI specification
type GetRoutingControlStateRequest struct {
	Routingcontrolarn interface{} `json:"RoutingControlArn"`
}

// UpdateRoutingControlStatesRequest represents the UpdateRoutingControlStatesRequest schema from the OpenAPI specification
type UpdateRoutingControlStatesRequest struct {
	Updateroutingcontrolstateentries interface{} `json:"UpdateRoutingControlStateEntries"`
	Safetyrulestooverride interface{} `json:"SafetyRulesToOverride,omitempty"`
}

// UpdateRoutingControlStateRequest represents the UpdateRoutingControlStateRequest schema from the OpenAPI specification
type UpdateRoutingControlStateRequest struct {
	Routingcontrolarn interface{} `json:"RoutingControlArn"`
	Routingcontrolstate interface{} `json:"RoutingControlState"`
	Safetyrulestooverride interface{} `json:"SafetyRulesToOverride,omitempty"`
}

// UpdateRoutingControlStatesResponse represents the UpdateRoutingControlStatesResponse schema from the OpenAPI specification
type UpdateRoutingControlStatesResponse struct {
}

// GetRoutingControlStateResponse represents the GetRoutingControlStateResponse schema from the OpenAPI specification
type GetRoutingControlStateResponse struct {
	Routingcontrolarn interface{} `json:"RoutingControlArn"`
	Routingcontrolname interface{} `json:"RoutingControlName,omitempty"`
	Routingcontrolstate interface{} `json:"RoutingControlState"`
}

// UpdateRoutingControlStateResponse represents the UpdateRoutingControlStateResponse schema from the OpenAPI specification
type UpdateRoutingControlStateResponse struct {
}

// ListRoutingControlsRequest represents the ListRoutingControlsRequest schema from the OpenAPI specification
type ListRoutingControlsRequest struct {
	Controlpanelarn interface{} `json:"ControlPanelArn,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// UpdateRoutingControlStateEntry represents the UpdateRoutingControlStateEntry schema from the OpenAPI specification
type UpdateRoutingControlStateEntry struct {
	Routingcontrolarn interface{} `json:"RoutingControlArn"`
	Routingcontrolstate interface{} `json:"RoutingControlState"`
}
