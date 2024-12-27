package gc

import (
	"github.com/goharbor/harbor-cli/pkg/api"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// StopGCCommand creates the stop GC command
func StopGCCommand() *cobra.Command {
	var gcID int64

	cmd := &cobra.Command{
		Use:     "stop",
		Short:   "Stop the garbage collection process",
		Long:    "Stop the currently running garbage collection process in Harbor using its GC ID",
		Example: "harbor-cli gc stop --gc-id 1234",
		RunE: func(cmd *cobra.Command, args []string) error {
			if gcID == 0 {
				logrus.Error("GC ID is required")
				return nil
			}

			logrus.Infof("Stopping the GC process with GC ID: %d", gcID)

			// Call API to stop the GC process with the given gcID
			err := api.StopGC(gcID)
			if err != nil {
				logrus.Errorf("Failed to stop GC with ID %d: %v", gcID, err)
				return err
			}

			logrus.Infof("GC process with ID %d stopped successfully", gcID)
			return nil
		},
	}

	// Adding a flag for the gcID
	cmd.Flags().Int64VarP(&gcID, "gc-id", "g", 0, "GC ID to stop the garbage collection process")

	return cmd
}
