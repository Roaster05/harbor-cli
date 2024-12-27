package gc

import (
	"github.com/spf13/cobra"
)

// GC returns the top-level 'gc' command which houses all GC-related subcommands
func GC() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gc",
		Short: "Manage Garbage Collection (GC) schedules in Harbor",
	}
	cmd.AddCommand(
		CreateGCScheduleCommand(),
		StopGCCommand(),
		UpdateGCScheduleCommand(),
		GetGCScheduleCommand(),
		GetGCHistoryCommand(),
		GetGCLogCommand(),
		getGCScheduleCmd(),
	)

	return cmd
}
