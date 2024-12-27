package gc

import (
	"fmt"

	"github.com/goharbor/harbor-cli/pkg/api"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// GetGCHistoryCommand retrieves the history of garbage collection executions
func GetGCHistoryCommand() *cobra.Command {
	var page int64
	var pageSize int64
	var query string
	var sort string

	cmd := &cobra.Command{
		Use:     "history",
		Short:   "Get the garbage collection history",
		Long:    "Retrieve the history of garbage collection executions in Harbor, including details about previous GC runs.",
		Example: "harbor gc history --page 1 --page-size 10 --query <query> --sort <sort>",
		Run: func(cmd *cobra.Command, args []string) {
			resp, err := api.GetGCHistory(page, pageSize, query, sort)
			if err != nil {
				log.Errorf("Failed to get GC history: %v", err)
				return
			}
			fmt.Printf(resp.Link)
			// Log the response or handle as necessary
			log.Infof("GC History retrieved successfully: %v", resp)
		},
	}

	// Adding flags for pagination, query, and sorting
	flags := cmd.Flags()
	flags.Int64VarP(&page, "page", "p", 1, "Page number for pagination")
	flags.Int64VarP(&pageSize, "page-size", "s", 10, "Number of items per page")
	flags.StringVarP(&query, "query", "q", "", "Query to filter the GC history")
	flags.StringVarP(&sort, "sort", "r", "", "Sort order for the results")

	return cmd
}
