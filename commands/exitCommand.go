package commands

import "os"

func exitCommand(_ *CliConfig) error {
	os.Exit(0)
	return nil
}
