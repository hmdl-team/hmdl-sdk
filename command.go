package sdk

import "github.com/spf13/cobra"

type Command struct {
	Use string
	Run func(service Service, cmd *cobra.Command, args []string)
}
