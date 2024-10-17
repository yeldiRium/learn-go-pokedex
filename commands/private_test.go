package commands

func (c *CliConfig) WithNextMapUrl(nextUrl string) *CliConfig {
	c.nextMapUrl = &nextUrl
	return c
}

func (c *CliConfig) WithPreviousMapUrl(previousUrl string) *CliConfig {
	c.previousMapUrl = &previousUrl
	return c
}
