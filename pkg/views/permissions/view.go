package permissions

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/permissions"
	"github.com/goharbor/harbor-cli/pkg/views/base/tablelist"
)

var columns = []table.Column{
	{Title: "Scope", Width: 15},
	{Title: "Action", Width: 20},
	{Title: "Resource", Width: 30},
}

// PrintPermissions displays project and system-level permissions in a tabular format.
func PrintPermissions(permissions *permissions.GetPermissionsOK) {
	if permissions == nil || permissions.Payload == nil {
		fmt.Println("No permissions available.")
		return
	}

	var rows []table.Row

	// Add project-level permissions to the table
	for _, perm := range permissions.Payload.Project {
		rows = append(rows, table.Row{
			"Project",
			perm.Action,
			perm.Resource,
		})
	}

	// Add system-level permissions to the table
	for _, perm := range permissions.Payload.System {
		rows = append(rows, table.Row{
			"System",
			perm.Action,
			perm.Resource,
		})
	}

	// Create a table model and render it
	m := tablelist.NewModel(columns, rows, len(rows))
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
