package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var Debug bool
var Quiet bool
var PrintLevel string
var SkipDirList []string

func init() {
	RootCmd.AddCommand(CmdAnsible)

	RootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "debug level logging")
	RootCmd.PersistentFlags().BoolVarP(&Quiet, "quiet", "q", false, "run with error level only logging")

	RootCmd.PersistentFlags().StringVar(&PrintLevel, "level", "file", "print level (file|role|project)")
	RootCmd.PersistentFlags().StringSliceVar(&SkipDirList, "skip-dirs", make([]string, 0), "comma separated list of directories to skip")
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
	Use:     "iac-count",
	Short:   "An analyzer for iac projects",
	Long:    `iac-count is an analyzer for iac projects.`,
	Version: "0.1",
	Run: func(cmd *cobra.Command, args []string) {
		configureLogging()
		err := cmd.Usage()
		if err != nil {
			log.Fatal(err)
		}
	},
}
