package commands

import (
	"github.com/buildpacks/pack/pkg/logging"
	"github.com/spf13/cobra"
)

// ManifestDeleteFlags define flags provided to the ManifestDelete
// type ManifestDeleteFlags struct {
// }

// ManifestDelete deletes one or more manifest lists from local storage
func ManifestDelete(logger logging.Logger, pack PackClient) *cobra.Command {
	// var flags ManifestDeleteFlags

	cmd := &cobra.Command{
		Use: "pack manifest remove [manifest-list] [manifest-list...] [flags]",
		Args: cobra.MatchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
		Short: "Delete one or more manifest lists from local storage",
		Example: `pack manifest remove cnbs/sample-package:hello-multiarch-universe`,
		Long: `Delete one or more manifest lists from local storage.

		When a manifest list exits locally, users can remove existing images from a manifest list`,
		RunE: logError(logger, func(cmd *cobra.Command, args []string) error {
			return nil
		}),
	}

	AddHelpFlag(cmd, "remove")
	return cmd
}