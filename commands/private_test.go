package commands

func (c *CliConfig) WithNextMapUrl(nextUrl string) *CliConfig {
	c.nextMapUrl = &nextUrl
	return c
}
