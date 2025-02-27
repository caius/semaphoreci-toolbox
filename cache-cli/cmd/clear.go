package cmd

import (
	"fmt"

	"github.com/semaphoreci/toolbox/cache-cli/pkg/storage"
	"github.com/semaphoreci/toolbox/cache-cli/pkg/utils"
	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Remove all keys in the cache.",
	Long:  ``,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		RunClear(cmd, args)
	},
}

func RunClear(cmd *cobra.Command, args []string) {
	storage, err := storage.InitStorage()
	utils.Check(err)

	err = storage.Clear()
	utils.Check(err)
	fmt.Println("Deleted all caches.")
}

func init() {
	RootCmd.AddCommand(clearCmd)
}
