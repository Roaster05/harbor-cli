package gc

import (
	"time"

	"github.com/goharbor/harbor-cli/pkg/api"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// UpdateGCScheduleCommand creates the update GC schedule command
func UpdateGCScheduleCommand() *cobra.Command {
	var scheduleType, cronExpr string
	var nextScheduledTime string

	cmd := &cobra.Command{
		Use:     "update",
		Short:   "Update the GC schedule",
		Long:    "Update the garbage collection schedule in Harbor",
		Example: "harbor-cli gc update --schedule-type <type> --cron <cron-expression> --next-scheduled-time <time>",
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			// Parse nextScheduledTime
			parsedTime, parseErr := time.Parse(time.RFC3339, nextScheduledTime)
			if parseErr != nil {
				logrus.Errorf("Failed to parse nextScheduledTime: %v", parseErr)
				return parseErr
			}

			// Call API to update the schedule
			err = api.UpdateGCSchedule(scheduleType, cronExpr, &parsedTime)
			if err != nil {
				logrus.Errorf("Failed to update GC schedule: %v", err)
				return err
			}

			logrus.Infof("GC schedule updated successfully")
			return nil
		},
	}

	cmd.Flags().StringVarP(&scheduleType, "schedule-type", "t", "", "Type of the GC schedule (e.g., daily, weekly)")
	cmd.Flags().StringVarP(&cronExpr, "cron", "", "", "Cron expression for the schedule")
	cmd.Flags().StringVarP(&nextScheduledTime, "next-scheduled-time", "n", "", "Next scheduled time for the GC run (RFC3339 format)")

	return cmd
}
