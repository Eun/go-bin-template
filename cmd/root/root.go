package root

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	gobintemplate "github.com/Eun/go-bin-template/pkg/go-bin-template"
)

var (
	rootCmd = &cobra.Command{
		RunE: run,
	}
	boolFlag bool
)

// Execute executes the root cmd.
func Execute() {
	rootCmd.Flags().BoolVarP(&boolFlag, "boolFlag", "b", false, "some bool flag")
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}

// Version sets the rootCmd version to the specified value.
func Version(version string) {
	rootCmd.Version = version
}

func run(cmd *cobra.Command, args []string) error {
	gobintemplate.HelloWorld(boolFlag)
	return nil
}
