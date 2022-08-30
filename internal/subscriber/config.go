package subscriber

type Config struct {
	ClusterID string `toml:"cluster_id"`
	ClientID  string `toml:"client_id"`
	Subject   string `toml:"subject"`
}

func NewConfig() *Config {
	return &Config{}
}
