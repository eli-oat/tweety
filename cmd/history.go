package cmd

import (
	"encoding/json"
	"os"

	"github.com/spf13/cobra"
)

func NewCmdHistorySearch() *cobra.Command {
	return &cobra.Command{
		Use:  "search",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			query := args[0]
			res, err := sendMessage[any](map[string]any{
				"command": "history.search",
				"query":   query,
			})
			if err != nil {
				return err
			}

			encoder := json.NewEncoder(os.Stdout)
			if err := encoder.Encode(res); err != nil {
				return err
			}

			return nil
		},
	}
}
func NewCmdHistory() *cobra.Command {
	cmd := &cobra.Command{
		Use: "history",
	}

	cmd.AddCommand(NewCmdHistorySearch())

	return cmd
}
