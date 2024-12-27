package gc

import (
	"time"

	"github.com/goharbor/harbor-cli/pkg/api"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// CreateGCScheduleCommand creates a new GC schedule in Harbor
func CreateGCScheduleCommand() *cobra.Command {
	var scheduleName string
	var cronExpression string
	var nextScheduledTime string // You can later parse this into a time.Time

	cmd := &cobra.Command{
		Use:     "create",
		Short:   "Create a new GC schedule",
		Long:    "Create a new Garbage Collection schedule in Harbor using a cron expression",
		Example: "harbor gc create --name <schedule_name> --cron <cron_expression> --next-scheduled-time <time>",
		Run: func(cmd *cobra.Command, args []string) {
			if scheduleName == "" || cronExpression == "" {
				log.Errorf("Schedule name and cron expression are required")
				return
			}

			// Parsing the next scheduled time if it's provided
			var parsedTime *time.Time
			if nextScheduledTime != "" {
				t, err := time.Parse(time.RFC3339, nextScheduledTime) // Assuming the time format is RFC3339
				if err != nil {
					log.Errorf("Failed to parse next scheduled time: %v", err)
					return
				}
				parsedTime = &t
			}

			// Call the API to create the GC schedule
			err := api.CreateGCSchedule(scheduleName, cronExpression, parsedTime)
			if err != nil {
				log.Errorf("Failed to create GC schedule: %v", err)
			} else {
				log.Infof("GC schedule '%s' created successfully with cron '%s'", scheduleName, cronExpression)
			}
		},
	}

	// Adding flags for schedule name, cron expression, and next scheduled time
	flags := cmd.Flags()
	flags.StringVarP(&scheduleName, "name", "n", "", "Name of the GC schedule")
	flags.StringVar(&cronExpression, "cron", "", "Cron expression for the GC schedule (no shorthand)")
	flags.StringVarP(&nextScheduledTime, "next-scheduled-time", "t", "", "Next scheduled time for the GC in RFC3339 format")

	return cmd
}
