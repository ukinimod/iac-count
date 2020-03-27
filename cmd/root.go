package cmd

import (
	"os"

	"github.com/MaibornWolff/iac-count/pkg/core"
	"github.com/MaibornWolff/iac-count/pkg/output"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var Debug bool
var Quiet bool
var PrintLevel string

func init() {
	RootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "debug level logging")
	RootCmd.PersistentFlags().BoolVarP(&Quiet, "quiet", "q", false, "error level logging only")

	RootCmd.PersistentFlags().StringVar(&PrintLevel, "level", "file", "print level (file|role|project)")
}

func configureLogging() {
	if Debug {
		log.SetLevel(log.DebugLevel)
	} else if Quiet {
		log.SetLevel(log.ErrorLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}

	log.SetOutput(os.Stderr)
}

var RootCmd = &cobra.Command{
	Use:   "ANSIBLE_ROOT",
	Short: "analyzes projects",
	Long:  "analyzes projects",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		configureLogging()

		fileMetrics := core.DirectoryCreator{}.CreateFromPath(args[0], nil).Analyze()
		output.PrintMetricsAsCsv(fileMetrics, PrintLevel)
	},
}
