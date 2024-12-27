package gc

import (
	"fmt"

	"github.com/goharbor/harbor-cli/pkg/api"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// GetGCScheduleCommand retrieves the current garbage collection schedule
func GetGCScheduleCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "schedule",
		Short:   "Get the current GC schedule",
		Long:    "Retrieve the current garbage collection schedule in Harbor.",
		Example: "harbor gc schedule",
		Run: func(cmd *cobra.Command, args []string) {
			resp, err := api.GetGCSchedule()
			if err != nil {
				log.Errorf("Failed to get GC schedule: %v", err)
				return
			}
			fmt.Printf(resp.Payload.JobStatus, resp.Payload.JobName, resp.Payload.JobKind, resp.Payload.JobParameters)
			// Log the response or handle as necessary
			log.Infof("GC Schedule retrieved successfully: %v", resp)
		},
	}

	return cmd
}
