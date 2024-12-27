package api

import (
	"fmt"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/gc"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
	"github.com/goharbor/harbor-cli/pkg/utils"
	log "github.com/sirupsen/logrus"
)

// CreateGCSchedule creates a new Garbage Collection schedule based on parameters.

func CreateGCSchedule(scheduleType, cronExpr string, nextScheduledTime *time.Time) error {
	ctx, client, err := utils.ContextWithClient()
	if err != nil {
		return err
	}

	// Dynamically create the ScheduleObj based on input
	scheduleObj := &models.ScheduleObj{
		Type: scheduleType,
	}

	// Handle schedule types and their required fields
	switch scheduleType {
	case "Schedule", "Hourly", "Daily", "Weekly", "Custom":
		if cronExpr == "" {
			log.Errorf("Cron expression is required for schedule type '%s'.", scheduleType)
			return fmt.Errorf("cron expression is required for schedule type '%s'", scheduleType)
		}
		scheduleObj.Cron = cronExpr

		if nextScheduledTime != nil {
			scheduleObj.NextScheduledTime = strfmt.DateTime(*nextScheduledTime)
		} else {
			scheduleObj.NextScheduledTime = strfmt.DateTime(time.Now().Add(24 * time.Hour))
		}

	case "Manual":
		scheduleObj.Cron = ""
		scheduleObj.NextScheduledTime = strfmt.DateTime(time.Now())

	case "None":
		scheduleObj.Cron = ""
		scheduleObj.NextScheduledTime = strfmt.DateTime{}

	default:
		log.Errorf("Invalid schedule type: '%s'", scheduleType)
		return fmt.Errorf("invalid schedule type: '%s'", scheduleType)
	}

	// API call to create the GC schedule
	_, err = client.GC.CreateGCSchedule(ctx, &gc.CreateGCScheduleParams{
		Schedule: &models.Schedule{
			ID:       100,
			Schedule: scheduleObj,
		},
	})
	if err != nil {
		log.Errorf("Failed to create GC schedule: %v", err)
		return err
	}

	log.Infof("GC schedule of type '%s' created successfully.", scheduleType)
	return nil
}

// GetGC retrieves the status of a specific Garbage Collection execution.
func GetGC(gcID int64) (*gc.GetGCOK, error) {
	ctx, client, err := utils.ContextWithClient()
	if err != nil {
		return nil, err
	}

	resp, err := client.GC.GetGC(ctx, &gc.GetGCParams{GCID: gcID})
	if err != nil {
		log.Errorf("Failed to get GC status: %v", err)
		return nil, err
	}

	return resp, nil
}

// GetGCHistory retrieves the Garbage Collection execution history.
func GetGCHistory(page, pageSize int64, query, sort string) (*gc.GetGCHistoryOK, error) {
	ctx, client, err := utils.ContextWithClient()
	if err != nil {
		return nil, err
	}

	resp, err := client.GC.GetGCHistory(ctx, &gc.GetGCHistoryParams{
		Page:     &page,
		PageSize: &pageSize,
		Q:        &query,
		Sort:     &sort,
	})
	if err != nil {
		log.Errorf("Failed to get GC history: %v", err)
		return nil, err
	}

	return resp, nil
}

// GetGCLog retrieves logs for a specific Garbage Collection execution.
func GetGCLog(gcID int64) (*gc.GetGCLogOK, error) {
	ctx, client, err := utils.ContextWithClient()
	if err != nil {
		return nil, err
	}

	resp, err := client.GC.GetGCLog(ctx, &gc.GetGCLogParams{GCID: gcID})
	if err != nil {
		log.Errorf("Failed to get GC logs: %v", err)
		return nil, err
	}

	return resp, nil
}

// GetGCSchedule retrieves the current Garbage Collection schedule.
func GetGCSchedule() (*gc.GetGCScheduleOK, error) {
	ctx, client, err := utils.ContextWithClient()
	if err != nil {
		return nil, err
	}

	resp, err := client.GC.GetGCSchedule(ctx, &gc.GetGCScheduleParams{})
	if err != nil {
		log.Errorf("Failed to get GC schedule: %v", err)
		return nil, err
	}

	return resp, nil
}

// StopGC stops a specific Garbage Collection execution.
func StopGC(gcID int64) error {
	ctx, client, err := utils.ContextWithClient()
	if err != nil {
		return err
	}

	_, err = client.GC.StopGC(ctx, &gc.StopGCParams{GCID: gcID})
	if err != nil {
		log.Errorf("Failed to stop GC: %v", err)
		return err
	}

	log.Infof("GC execution stopped successfully: ID=%d", gcID)
	return nil
}

// UpdateGCSchedule updates an existing Garbage Collection schedule based on parameters.
func UpdateGCSchedule(scheduleType, cronExpr string, nextScheduledTime *time.Time) error {
	ctx, client, err := utils.ContextWithClient()
	if err != nil {
		return err
	}

	// Dynamically create the ScheduleObj based on input
	scheduleObj := &models.ScheduleObj{
		Type: scheduleType,
	}

	// Handle schedule types and their required fields
	switch scheduleType {
	case "Schedule", "Hourly", "Daily", "Weekly", "Custom":
		if cronExpr == "" {
			log.Errorf("Cron expression is required for schedule type '%s'.", scheduleType)
			return fmt.Errorf("cron expression is required for schedule type '%s'", scheduleType)
		}
		scheduleObj.Cron = cronExpr

		if nextScheduledTime != nil {
			scheduleObj.NextScheduledTime = strfmt.DateTime(*nextScheduledTime)
		} else {
			scheduleObj.NextScheduledTime = strfmt.DateTime(time.Now().Add(24 * time.Hour))
		}

	case "Manual":
		scheduleObj.Cron = ""
		scheduleObj.NextScheduledTime = strfmt.DateTime(time.Now())

	case "None":
		scheduleObj.Cron = ""
		scheduleObj.NextScheduledTime = strfmt.DateTime{}

	default:
		log.Errorf("Invalid schedule type: '%s'", scheduleType)
		return fmt.Errorf("invalid schedule type: '%s'", scheduleType)
	}

	// API call to update the GC schedule
	_, err = client.GC.UpdateGCSchedule(ctx, &gc.UpdateGCScheduleParams{
		Schedule: &models.Schedule{
			Schedule: scheduleObj,
		},
	})
	if err != nil {
		log.Errorf("Failed to update GC schedule: %v", err)
		return err
	}

	log.Infof("GC schedule of type '%s' updated successfully.", scheduleType)
	return nil
}
