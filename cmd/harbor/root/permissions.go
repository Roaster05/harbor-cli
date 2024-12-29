package root

import (
	"github.com/goharbor/harbor-cli/pkg/api"
	"github.com/goharbor/harbor-cli/pkg/views/permissions"
	"github.com/spf13/cobra"
)

func PermissionsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "permissions",
		Short: "Get the permissions for the User",
		RunE: func(cmd *cobra.Command, args []string) error {
			permission, err := api.GetPermissions()
			if err != nil {
				return err
			}
			permissions.PrintPermissions(permission)
			return nil
		},
		Example: `  # Get the permissions for the User`,
	}

	return cmd

}
