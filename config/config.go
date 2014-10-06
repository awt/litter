package config

type envMap map[string]string

type Config struct {
	Name string
	values map[string]envMap
}

func (c *Config) SetEnvironment(env string) {
	c.Name = env
	c.values = make(map[string]envMap)
}

func (c *Config) Get(key string) (value string) {
	if "" == c.Name {
		c.Name = "local"
	}

	return c.values[c.Name][key]
}

func (c *Config) Set(key string, value string) {
	if c.values[c.Name] == nil {
		c.values[c.Name] = envMap{}
	}
	c.values[c.Name][key] = value
}
