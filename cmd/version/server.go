package version

import (
	"fmt"
	"github.com/link1st/gowebsocket/common/global"
	"github.com/spf13/cobra"
)

var (
	StartCmd = &cobra.Command{
		Use:     "version",
		Short:   "Get version info",
		Example: "go-chat version",
		PreRun: func(cmd *cobra.Command, args []string) {

		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func run() error {
	fmt.Println(global.Version)
	return nil
}
