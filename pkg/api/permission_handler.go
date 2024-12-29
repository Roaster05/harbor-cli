package api

import (
	"fmt"

	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/permissions"
	"github.com/goharbor/harbor-cli/pkg/utils"
)

func GetPermissions() (*permissions.GetPermissionsOK, error) {
	ctx, client, err := utils.ContextWithClient()
	if err != nil {
		return nil, err
	}

	response, err := client.Permissions.GetPermissions(ctx, &permissions.GetPermissionsParams{})
	if err != nil {
		return nil, fmt.Errorf("error getting permissions: %w", err)
	}

	return response, nil
}
