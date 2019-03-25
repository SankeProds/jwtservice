package implementation

type HardCodedConfig struct {
	host string
	port int
}

func (c *HardCodedConfig) GetPort() int {
	return c.port
}

func (c *HardCodedConfig) GetHost() string {
	return c.host
}

func (c *HardCodedConfig) Load() {
	c.host = ""
	c.port = 1234
}
