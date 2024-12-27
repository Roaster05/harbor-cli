package gc

import (
	"github.com/goharbor/harbor-cli/pkg/api"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// GetGCLogCommand retrieves the logs for a specific Garbage Collection execution
func GetGCLogCommand() *cobra.Command {
	var gcID int64

	cmd := &cobra.Command{
		Use:     "log",
		Short:   "Get the logs of a specific GC execution",
		Long:    "Retrieve the logs of a specific garbage collection execution in Harbor by providing the GC ID.",
		Example: "harbor gc log --gc-id <gc_id>",
		Run: func(cmd *cobra.Command, args []string) {
			resp, err := api.GetGCLog(gcID)
			if err != nil {
				log.Errorf("Failed to get GC logs: %v", err)
				return
			}

			// Log the response or handle as necessary
			log.Infof("GC Log retrieved successfully: %v", resp)
		},
	}

	// Adding flag for GC ID
	flags := cmd.Flags()
	flags.Int64VarP(&gcID, "gc-id", "g", 0, "GC ID to retrieve logs for")
	return cmd
}
