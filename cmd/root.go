package cmd

import (
	"log"
	"os"

	"github.com/MaibornWolff/iac-count/pkg/core"
	"github.com/hashicorp/logutils"
	"github.com/spf13/cobra"
)

var Debug bool
var Quiet bool
var PrintLevel string
var MetricList []string

func init() {
	RootCmd.AddCommand(CmdAnsible)
	RootCmd.AddCommand(CmdDescribe)

	RootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "debug level logging")
	RootCmd.PersistentFlags().BoolVarP(&Quiet, "quiet", "q", false, "run with error level only logging")

	RootCmd.PersistentFlags().StringVar(&PrintLevel, "level", "file", "print level (file|role|project)")
	RootCmd.PersistentFlags().StringSliceVarP(&MetricList, "metrics", "m", core.MetricNames[:], "comma separated list of metrics (default: all)")
}

func configureLogging() {
	var logLevel = logutils.LogLevel("INFO")
	if Debug {
		logLevel = logutils.LogLevel("DEBUG")
	}
	if Quiet {
		logLevel = logutils.LogLevel("ERROR")
	}

	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "INFO", "WARN", "ERROR"},
		MinLevel: logLevel,
		Writer:   os.Stderr,
	}
	log.SetOutput(filter)
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
