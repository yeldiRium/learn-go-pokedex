package commands

import "os"

func exitCommand(_ *CliConfig, _ []string) error {
	os.Exit(0)
	return nil
}
