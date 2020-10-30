package cmd

type FilterConfig struct {
	Tags          []map[string]interface{}
	ResourceTypes []string
}

var Filter FilterConfig

func (f *FilterConfig) get() FilterConfig {
	return *f
}
