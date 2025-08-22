package main

import (
	"github.com/route53-recovery-cluster/mcp-server/config"
	"github.com/route53-recovery-cluster/mcp-server/models"
	tools_x_amz_target_togglecustomerapi_listroutingcontrols "github.com/route53-recovery-cluster/mcp-server/tools/x_amz_target_togglecustomerapi_listroutingcontrols"
	tools_x_amz_target_togglecustomerapi_updateroutingcontrolstate "github.com/route53-recovery-cluster/mcp-server/tools/x_amz_target_togglecustomerapi_updateroutingcontrolstate"
	tools_x_amz_target_togglecustomerapi_updateroutingcontrolstates "github.com/route53-recovery-cluster/mcp-server/tools/x_amz_target_togglecustomerapi_updateroutingcontrolstates"
	tools_x_amz_target_togglecustomerapi_getroutingcontrolstate "github.com/route53-recovery-cluster/mcp-server/tools/x_amz_target_togglecustomerapi_getroutingcontrolstate"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_x_amz_target_togglecustomerapi_listroutingcontrols.CreateListroutingcontrolsTool(cfg),
		tools_x_amz_target_togglecustomerapi_updateroutingcontrolstate.CreateUpdateroutingcontrolstateTool(cfg),
		tools_x_amz_target_togglecustomerapi_updateroutingcontrolstates.CreateUpdateroutingcontrolstatesTool(cfg),
		tools_x_amz_target_togglecustomerapi_getroutingcontrolstate.CreateGetroutingcontrolstateTool(cfg),
	}
}
