package config

type Config struct {
	Rules []RuleConfig `yaml:"rules"`
}

type RuleConfig struct {
	Name       string       `yaml:"name"`
	Match      MatchConfig  `yaml:"match"`
	Percentage float64      `yaml:"percentage"`
	Action     ActionConfig `yaml:"action"`
}

type MatchConfig struct {
	Host       string `yaml:"host"`
	PathPrefix string `yaml:"path_prefix"`
	Method     string `yaml:"method"`
}

type ActionConfig struct {
	Type   string            `yaml:"type"`
	Target string            `yaml:"target"`
	Status int               `yaml:"status"`
	Body   string            `yaml:"body"`
	Delay  int               `yaml:"delay_ms"`
	Header map[string]string `yaml:"headers"`
}
